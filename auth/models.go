package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

//TODO:Заменить на чтение из переменных окружения
const AccessSecretKey = "fsd23fsdf24fdwsf"
const RefreshSecretKey = "sdfsd344gdfgs34dfawe4r34dfgf4d1vbaf42fsda"

type Claims struct {
	jwt.StandardClaims
	Username    string
	AccessToken string
}

func NewAccessClaims() Claims {
	return Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
}

func NewRefreshClaims(username string, accessToken string) Claims {
	return Claims{
		Username:    username,
		AccessToken: accessToken,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
}

type DbRefreshToken struct {
	Id       uint64
	Username string
	Token    string
}

func (DbRefreshToken) TableName() string {
	return "refresh_tokens"
}

type User struct {
	Id          int
	Username    string
	Password    string
	Description string
	Active      bool
}

func (User) TableName() string {
	return "users"
}
