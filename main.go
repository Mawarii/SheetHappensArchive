package main

import (
	"os"
	"sheethappens/backend/controller"
	"sheethappens/backend/database"
	"sheethappens/backend/handler"

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

	restriced := e.Group("/main")
	restriced.Use(handler.CheckSessionToken)

	// character
	restriced.GET("/character", func(c echo.Context) error {
		return c.File("frontend/html/index.html")
	}).Name = "character"
	restriced.POST("/character", controller.CreateCharacter)
	restriced.GET("/character/:id", controller.GetCharacter)
	restriced.PUT("/character/:id", controller.UpdateCharacter)
	restriced.DELETE("/character/:id", controller.DeleteCharacter)
	restriced.GET("/character/list", controller.GetAllCharacters)

	e.Logger.Fatal(e.Start(":8080"))
}
