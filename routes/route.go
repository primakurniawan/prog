package routes

import (
	"prog/controllers"
	"prog/middlewares"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controllers.GetAllUsers)
	e.POST("/users", controllers.CreateUser)

	middlewares.Logger(e)
	eAuthBasic := e.Group("/auth")
	eAuthBasic.Use(middleware.BasicAuth(middlewares.BasicAuth))
	eAuthBasic.GET("/users", controllers.GetAllUsers)

	return e

}
