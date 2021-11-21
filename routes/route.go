package routes

import (
	"prog/constants"
	"prog/controllers"
	"prog/middlewares"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controllers.GetAllUsers)
	e.POST("/users", controllers.CreateUser)
	e.POST("/login", controllers.Login)

	middlewares.Logger(e)
	eAuthBasic := e.Group("/auth")
	eAuthBasic.Use(middleware.BasicAuth(middlewares.BasicAuth))
	eAuthBasic.GET("/users", controllers.GetAllUsers)

	eJWT := e.Group("/jwt")
	eJWT.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	eJWT.GET("/users", controllers.GetAllUsers)

	return e

}
