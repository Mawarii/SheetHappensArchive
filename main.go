package main

import (
	"net/http"
	"os"
	"sheethappens/backend/controller"
	"sheethappens/backend/database"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	sessionSecret := os.Getenv("SECRET")

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(sessionSecret))))

	e.Static("/static", "frontend")

	database.DatabaseInit()
	gorm := database.DB()

	dbGorm, err := gorm.DB()
	dbGorm.Ping()

	if err != nil {
		panic(err)
	}

	e.GET("/", controller.Login)
	e.POST("/", controller.Authenticate)

	e.GET("/register", controller.Register)
	e.POST("/register", controller.RegisterUser)

	restriced := e.Group("/character")
	restriced.Use(checkSessionToken)

	restriced.GET("", func(c echo.Context) error {
		return c.File("frontend/html/index.html")
	}).Name = "character"
	restriced.POST("", controller.GetAllCharacters)
	restriced.GET("/:id", controller.GetCharacter)
	restriced.PUT("/:id", controller.UpdateCharacter)
	restriced.DELETE("/:id", controller.DeleteCharacter)

	e.GET("/characters", controller.GetAllCharacters)

	e.Logger.Fatal(e.Start(":8080"))
}

func checkSessionToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		username, ok := sess.Values["username"].(string)

		if ok && username != "" {
			return next(c)
		}

		return c.Redirect(http.StatusSeeOther, "/")
	}
}
