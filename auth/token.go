package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

func IsValidAccessToken(AccessToken string) (bool, error) {
	token, err := jwt.ParseWithClaims(AccessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(AccessSecretKey), nil
	})

	if err != nil {
		return false, err
	}

	_, ok := token.Claims.(*Claims)

	if ok && token.Valid {
		return true, nil
	}

	return false, nil
}

func GetUsernameFormToken(RefreshToken string) (string, error) {
	token, err := jwt.ParseWithClaims(RefreshToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(RefreshSecretKey), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*Claims)

	if ok && claims.Username != "" {
		return claims.Username, nil
	}

	return "", errors.New("can't get username from token")
}

func IsValidAccessAndRefreshToken(AccessToken string, RefreshToken string) (bool, error) {
	token, err := jwt.ParseWithClaims(RefreshToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(RefreshSecretKey), nil
	})

	if err != nil {
		return false, err
	}

	claims, ok := token.Claims.(*Claims)

	if !ok {
		return false, errors.New("not ok while parse claims")
	} else if !token.Valid {
		return false, errors.New("invalid token")
	} else if claims.AccessToken != AccessToken {
		return false, errors.New("access token is different")
	}

	username, err := GetUsernameByToken(RefreshToken)

	if err != nil {
		return false, err
	}

	if username != claims.Username {
		return false, errors.New("username is different")
	}

	return true, nil
}

func CreateTokens(username string) (AccessToken string, RefreshToken string, err error) {
	atClaims := NewAccessClaims()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	AccessToken, err = at.SignedString([]byte(AccessSecretKey))

	if err != nil {
		return "", "", err
	}

	rtClaims := NewRefreshClaims(username, AccessToken)

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	RefreshToken, err = rt.SignedString([]byte(RefreshSecretKey))

	if err != nil {
		return "", "", err
	}

	err = RemoveTokenByUsername(username)

	if err != nil {
		return "", "", err
	}

	err = AddToken(RefreshToken, username)

	if err != nil {
		return "", "", err
	}

	return AccessToken, RefreshToken, nil
}
