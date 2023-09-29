package main

import (
	"sheethappens/backend/controller"
	"sheethappens/backend/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "frontend")

	e.GET("/", func(c echo.Context) error {
		return c.File("frontend/html/index.html")
	})

	database.DatabaseInit()
	gorm := database.DB()

	dbGorm, err := gorm.DB()
	dbGorm.Ping()

	if err != nil {
		panic(err)
	}

	e.GET("/register", controller.Register)
	e.POST("/register", controller.RegisterUser)

	e.GET("/login", controller.Login)
	e.POST("/login", controller.Authenticate)

	characterRoute := e.Group("/character")
	characterRoute.POST("", controller.CreateCharacter)
	characterRoute.GET("/:id", controller.GetCharacter)
	characterRoute.PUT("/:id", controller.UpdateCharacter)
	characterRoute.DELETE("/:id", controller.DeleteCharacter)

	e.GET("/characters", controller.GetAllCharacters)

	e.Logger.Fatal(e.Start(":8080"))
}
