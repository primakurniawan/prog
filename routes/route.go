package routes

import (
	"prog/factory"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	n := echo.New()
	e := n.Group("v1")

	userPresentation := factory.InitUser()
	e.POST("/users", userPresentation.UserHandler.RegisterUserHandler)
	e.GET("/users", userPresentation.UserHandler.GetAllUsersHandler)
	e.GET("/users/:userId", userPresentation.UserHandler.GetUserByIdHandler)
	e.GET("/users/:userId/following", userPresentation.UserHandler.GetUserFollowingByIdHandler)
	e.GET("/users/:userId/followers", userPresentation.UserHandler.GetUserFollowersByIdHandler)

	// middlewares.Logger(e)
	return n

}
