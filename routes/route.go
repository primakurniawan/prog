package routes

import (
	"prog/factory"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	userPresentation := factory.InitUser()
	e.POST("/users", userPresentation.UserHandler.RegisterUserHandler)
	e.GET("/users", userPresentation.UserHandler.GetAllUserHandler)
	e.GET("/users/:userId", userPresentation.UserHandler.GetUserByIdHandler)

	// middlewares.Logger(e)
	return e

}
