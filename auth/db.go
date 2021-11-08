package auth

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
)

var db *gorm.DB

func connectDB(user string, password string, host string, port string, database string) error {
	dbURi := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, database, port)

	//dialect := postgres.Open(dbURi)
	dialect := postgres.New(postgres.Config{
		DSN: dbURi,
	})

	conn, err := gorm.Open(dialect, &gorm.Config{})

	if err != nil {
		return err
	}
	db = conn
	return nil
}

func GetUser(username string) (*User, error) {
	user := User{}

	err := db.Where("LOWER(username) = ?", strings.ToLower(username)).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUsernameByToken(refreshToken string) (string, error) {
	token := DbRefreshToken{}

	err := db.Where("token = ?", refreshToken).Last(&token).Error

	if err != nil {
		return "", err
	}

	return token.Username, nil
}

func AddToken(refreshToken string, username string) error {
	token := DbRefreshToken{
		Token:    refreshToken,
		Username: username,
	}

	err := db.Select("username", "token").Create(&token).Error

	return err
}

func RemoveToken(refreshToken string) error {
	query := db.Where("token = ?", refreshToken).Delete(DbRefreshToken{})

	err := query.Error

	return err
}

func RemoveTokenByUsername(username string) error {
	query := db.Where("username = ?", username).Delete(DbRefreshToken{})

	err := query.Error

	return err
}
