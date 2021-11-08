package service

import (
	"Athena/auth"
	"Athena/auth/api"
	. "Athena/log"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
)

type GrpcService struct {
	api.UnimplementedAuthorizationServer
}

func (gs GrpcService) Authentication(c context.Context, user *api.User) (*api.TokenPair, error) {
	dbUser, err := auth.GetUser(user.Username)

	if err != nil {
		Info.Printf("User %s not found in db\n", user.Username)
		return nil, err
	}

	hasher := sha256.New()
	hasher.Write([]byte(user.Password))

	hash := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	if dbUser.Password != hash {
		Info.Printf("Wrond password from %s", user.Username)
		return nil, errors.New("wrong password")
	}

	if !dbUser.Active {
		Info.Printf("Inactive user %s", user.Username)
		return nil, errors.New("user is inactive")
	}

	at, rt, err := auth.CreateTokens(user.Username)

	if err != nil {
		Error.Printf("Failed to create token for %s. Error:%s\n", user.Username, err.Error())
		return nil, err
	}

	Info.Printf("Authentication success for %s", user.Username)
	return &api.TokenPair{
		AccessToken:  at,
		RefreshToken: rt,
	}, nil
}

func (gs GrpcService) Authorization(c context.Context, accessToken *api.AccessToken) (*api.Access, error) {
	ok, err := auth.IsValidAccessToken(accessToken.Token)

	if !ok {
		Info.Println("Invalid or expired access token. Access denied.")
		return &api.Access{Status: false}, err
	} else if err != nil {
		Info.Println("Error while parsing access token. Access denied.")
		return nil, err
	}

	Info.Println("Valid access token. Access granted.")
	return &api.Access{Status: true}, nil
}

func (gs GrpcService) RefreshTokens(c context.Context, tokensPair *api.TokenPair) (*api.TokenPair, error) {
	ok, err := auth.IsValidAccessAndRefreshToken(tokensPair.AccessToken, tokensPair.RefreshToken)

	if !ok {
		Info.Println("Invalid or expired refresh token. Access denied.")
		return nil, errors.New("tokens is not valid")
	} else if err != nil {
		Info.Println("Error while parsing refresh token. Access denied.")
		return nil, err
	}

	username, err := auth.GetUsernameFormToken(tokensPair.RefreshToken)

	if err != nil {
		Info.Println("Can't get username from refresh token. Access denied.")
		return nil, err
	}

	accessToken, refreshToken, err := auth.CreateTokens(username)

	if err != nil {
		Error.Printf("Failed to create tokens pair for %s. Access denied.\n", username)
		return nil, err
	}

	Info.Printf("Tokens pair refreshed for %s. Access granted.\n", username)
	return &api.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (gs GrpcService) Logout(c context.Context, tokenPair *api.TokenPair) (*api.EmptyResponse, error) {
	refreshToken := tokenPair.RefreshToken

	err := auth.RemoveToken(refreshToken)

	Info.Println("Manual logout")

	return &api.EmptyResponse{}, err
}
