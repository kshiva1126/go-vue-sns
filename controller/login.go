package controller

import (
	"go-vue-sns/db"
	"go-vue-sns/helper"
	"net/http"

	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
)

type Login struct {
	Email    string
	Password string
}

type LoginComplete struct {
	Success bool `json:"success"`
}

func LoginStart() echo.HandlerFunc {
	return func(c echo.Context) error {
		login := &Login{}
		if err := c.Bind(login); err != nil {
			return err
		}
		user := db.User{}
		db := db.GetConnection()
		db.Select("id, password").Where("email = ?", login.Email).Find(&user)
		err := helper.VerifyPassword(user.Password, login.Password)
		if err != nil {
			return c.String(http.StatusOK, "Fail to authenticate.")
		}

		// store data on session
		session := session.Default(c)
		session.Set("loginCompleted", "completed")
		session.Set("userId", user.ID)
		session.Save()

		return c.String(http.StatusOK, "Success to authenticate.")
	}
}

func SetLoginTemplate(c echo.Context) error {
	if helper.CheckLogin(c) {
		loginComplete := LoginComplete{
			Success: true,
		}

		return c.JSON(http.StatusOK, loginComplete)
	}
	return c.Render(http.StatusOK, "login", Login{})
}
