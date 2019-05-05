package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

func SetSignUpTemplate(c echo.Context) error {
	return c.Render(http.StatusOK, "signup", User{})
}
