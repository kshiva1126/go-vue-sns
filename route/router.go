package route

import (
	"go-vue-sns/controller"
	"io"
	"text/template"

	session "github.com/ipfans/echo-session"

	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func Init() *echo.Echo {
	e := echo.New()

	// set template
	t := &Template{
		templates: template.Must(template.ParseGlob("template/*.html")),
	}

	e.Renderer = t

	// set session
	store := session.NewCookieStore([]byte("secret-key"))
	store.MaxAge(86400)
	e.Use(session.Sessions("ESESSION", store))

	message := e.Group("/message")
	{
		message.GET("", controller.SetMessageTemplate)
		message.POST("", controller.CreateMessage())
		message.GET("/:mention_id", controller.GetMentions())
	}

	user := e.Group("/users")
	{
		user.GET("", controller.SetUserTemplate)
		user.GET("/:user_id", controller.GetUserDetail())
	}

	e.GET("/login", controller.SetLoginTemplate)
	e.POST("/login", controller.LoginStart())

	e.GET("/signup", controller.SetSignUpTemplate)
	e.POST("/signup", controller.CreateUser())
	return e
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
