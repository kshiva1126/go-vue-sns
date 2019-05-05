package controller

import (
	"fmt"
	"go-vue-sns/db"
	"go-vue-sns/helper"
	"net/http"

	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo"
)

type User struct {
	Name     string `json:"name"`
	Profile  string `json:"profile"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserInfo struct {
	Id       int       `json:"user_id"`
	Name     string    `json:"name"`
	Profile  string    `json:"profile"`
	Messages []Message `json:"messages"`
}

func SetUserTemplate(c echo.Context) error {
	// ログイン中のユーザ情報とユーザIDに紐づくメッセージ一覧を取得する
	userId, err := helper.GetUserId(c)
	if err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("%v", err))
	}
	db := db.GetConnection()
	defer db.Close()
	query := db.Table("users").
		Select("users.id, users.name, users.profile, messages.id as message_id, messages.body as message_body").
		Joins("LEFT JOIN messages ON users.id = messages.user_id").
		Where("users.id = ?", userId)
	rows, err := query.Rows()
	if err != nil {
		fmt.Println(err)
	}
	var UserInfoWithMessage struct {
		UserInfo
		MessageId   *string
		MessageBody *string
	}
	var messages []Message
	for rows.Next() {
		err := query.ScanRows(rows, &UserInfoWithMessage)
		if err != nil {
			fmt.Println(err)
		}
		if UserInfoWithMessage.MessageId != nil && UserInfoWithMessage.MessageBody != nil {
			message := Message{
				Id:   *UserInfoWithMessage.MessageId,
				Body: *UserInfoWithMessage.MessageBody,
			}
			messages = append(messages, message)
		}
	}

	userInfo := UserInfoWithMessage.UserInfo
	userInfo.Messages = messages
	fmt.Println(userInfo)
	return c.Render(http.StatusOK, "user", userInfo)
}

func GetUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		users := []db.User{}
		db := db.GetConnection()
		defer db.Close()
		data := db.Find(&users)
		return c.JSON(http.StatusOK, data)
	}
}

func GetUserDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		users := []db.User{}
		userId := c.Param("user_id")
		db := db.GetConnection()
		defer db.Close()
		data := db.Where("id = ?", userId).Find(&users)
		return c.JSON(http.StatusOK, data)
	}
}

func CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		u := &User{}
		if err := c.Bind(u); err != nil {
			return err
		}
		hash, err := helper.HashPassword(u.Password)
		if err != nil {
			panic(err)
		}
		user := db.User{
			Name:     u.Name,
			Profile:  u.Profile,
			Email:    u.Email,
			Password: hash,
		}
		db := db.GetConnection()
		defer db.Close()
		// create new user
		db.Create(&user)

		// get new userID
		db.Select("id").Where("email = ?", user.Email).Find(&user)

		// store data on session
		session := session.Default(c)
		session.Set("loginCompleted", "completed")
		session.Set("userId", user.ID)
		session.Save()

		return c.String(http.StatusOK, "Welcome to the brilliant world!")
	}
}
