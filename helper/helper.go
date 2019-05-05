package helper

import (
	"fmt"

	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func VerifyPassword(hash, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}

func CheckLogin(c echo.Context) bool {
	session := session.Default(c)
	loginId := session.Get("loginCompleted")
	if loginId != nil && loginId == "completed" {
		return true
	}

	return false
}

func GetUserId(c echo.Context) (userId uint, err error) {
	session := session.Default(c)
	sessUserId := session.Get("userId")
	if userId, ok := sessUserId.(uint); ok {
		return userId, nil
	}
	return userId, fmt.Errorf("Cannot get userid")
}
