package routes

import (
	"prog/constants"
	"prog/factory"
	"prog/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	articlePresentation := factory.InitArticle()
	eArticles := e.Group("/articles")
	configJWT := middleware.JWTConfig{
		SigningKey: []byte(constants.ACCESS_TOKEN_KEY),
		Claims:     &middlewares.JwtCustomClaims{},
	}
	eArticles.Use(middleware.JWTWithConfig(configJWT))
	eArticles.POST("", articlePresentation.ArticleHandler.CreateArticleHandler)
	eArticles.GET("", articlePresentation.ArticleHandler.GetAllArticleHandler)
	eArticles.GET("/:articleId", articlePresentation.ArticleHandler.GetArticleByIdHandler)
	eArticles.PATCH("/:articleId", articlePresentation.ArticleHandler.GetArticleByIdHandler)
	eArticles.DELETE("/:articleId", articlePresentation.ArticleHandler.GetArticleByIdHandler)
	// middlewares.Logger(e)
	return n

}
