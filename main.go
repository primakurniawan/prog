package main

import (
	"prog/controllers"
	"prog/db"

	"github.com/labstack/echo/v4"
)

func main() {
	db.InitDB()
	e := echo.New()

	e.GET("/users", controllers.GetAllUsers)
	e.POST("/users", controllers.CreateUser)

	e.Start(":8000")
}
