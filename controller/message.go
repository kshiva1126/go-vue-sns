package controller

import (
	"fmt"
	"go-vue-sns/db"
	"go-vue-sns/helper"
	"net/http"

	"github.com/labstack/echo"
)

type Message struct {
	Id   string `json:"message_id"`
	Body string `json:"body"`
}

func SetMessageTemplate(c echo.Context) error {
	if helper.CheckLogin(c) {
		return c.Render(http.StatusOK, "message", Message{})
	}

	return c.Render(http.StatusOK, "login", Login{})
}

func GetMessages() echo.HandlerFunc {
	return func(c echo.Context) error {
		messages := []db.Message{}
		db := db.GetConnection()
		defer db.Close()
		data := db.Find(&messages)
		return c.JSON(http.StatusOK, data)
	}
}

func GetMentions() echo.HandlerFunc {
	return func(c echo.Context) error {
		messages := []db.Message{}
		mentionId := c.Param("mention_id")
		db := db.GetConnection()
		defer db.Close()
		data := db.Where("mention_id = ?", mentionId).Find(&messages)
		return c.JSON(http.StatusOK, data)
	}
}

func CreateMessage() echo.HandlerFunc {
	return func(c echo.Context) error {
		if !helper.CheckLogin(c) {
			return c.Render(http.StatusOK, "login", Login{})
		}

		m := &Message{}
		if err := c.Bind(m); err != nil {
			return err
		}

		userId, err := helper.GetUserId(c)
		if err != nil {
			return c.String(http.StatusOK, fmt.Sprintf("%v", err))
		}
		message := db.Message{
			UserId: userId,
			Body:   m.Body,
		}
		db := db.GetConnection()
		db.Create(&message)
		return c.String(http.StatusOK, "Success to send message!")
	}
}
