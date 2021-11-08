package handler

import (
	"Athena/auth/api"
	"Athena/http_server/gRPC/auth"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authentication(c *gin.Context) {
	username := c.PostForm("username")

	password := c.PostForm("password")

	tokenPair, err := auth.Client.Authentication(context.Background(), &api.User{
		Username: username,
		Password: password,
	})

	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"message": "Неправильные логин или пароль",
		})
		return
	}

	accessToken := tokenPair.AccessToken
	refreshToken := tokenPair.RefreshToken
	SetCockies(c, accessToken, refreshToken, username)
	c.Redirect(http.StatusSeeOther, "/")
}

func AuthorizationMiddleware(c *gin.Context) {
	accessToken, err := c.Cookie("accessToken")

	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
		return
	}

	username, err := c.Cookie("username")

	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
		return
	}

	access, err := auth.Client.Authorization(context.Background(), &api.AccessToken{
		Token: accessToken,
	})

	if err != nil || !access.Status {
		refreshToken, err := c.Cookie("refreshToken")

		if err != nil {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		tokenPair, err := auth.Client.RefreshTokens(context.Background(), &api.TokenPair{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		})

		if err != nil {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		at := tokenPair.AccessToken
		rt := tokenPair.RefreshToken
		SetCockies(c, at, rt, username)
	}

	c.Next()
}

func Logout(c *gin.Context) {
	refreshToken, err := c.Cookie("refreshToken")

	if err != nil {
		c.Error(err)
	}

	accessToken, err := c.Cookie("accessToken")

	if err != nil {
		c.Error(err)
	}

	_, err = auth.Client.Logout(context.Background(), &api.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})

	if err != nil {
		c.Error(err)
	}

	SetCockies(c, "", "", "")

	c.Redirect(http.StatusSeeOther, "/login")
}

func SetCockies(c *gin.Context, accessToken, refreshToken, username string) {
	c.SetCookie("accessToken", accessToken, 3600*24*7, "/view",
		Domain, true, true)
	c.SetCookie("refreshToken", refreshToken, 3600*24*7, "/view",
		Domain, true, true)
	c.SetCookie("accessToken", accessToken, 3600*24*7, "/private",
		Domain, true, true)
	c.SetCookie("refreshToken", refreshToken, 3600*24*7, "/private",
		Domain, true, true)
	c.SetCookie("username", username, 3600*24*7, "/",
		Domain, true, true)
}
