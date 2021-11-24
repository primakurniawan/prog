package routes

import (
	"prog/factory"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	n := echo.New()
	e := n.Group("/v1")
	// config := middleware.JWTConfig{
	// 	Claims:     &middlewares.JwtCustomClaims{},
	// 	SigningKey: []byte(constants.ACCESS_TOKEN_KEY),
	// }
	// e.Use(middleware.JWTWithConfig(config))
	authPresentation := factory.InitAuth()
	eAuth := e.Group("/auth")
	eAuth.POST("", authPresentation.AuthHandler.LoginHandler)
	eAuth.PUT("", authPresentation.AuthHandler.ReLoginHandler)
	eAuth.DELETE("", authPresentation.AuthHandler.LogoutHandler)

	userPresentation := factory.InitUser()
	eUsers := e.Group("/users")
	eUsers.POST("", userPresentation.UserHandler.RegisterUserHandler)
	eUsers.GET("", userPresentation.UserHandler.GetAllUsersHandler)
	eUsers.GET("/:userId", userPresentation.UserHandler.GetUserByIdHandler)
	eUsers.GET("/:userId/following", userPresentation.UserHandler.GetUserFollowingByIdHandler)
	eUsers.GET("/:userId/followers", userPresentation.UserHandler.GetUserFollowersByIdHandler)

	// middlewares.Logger(e)
	return n

}
