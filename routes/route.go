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

	presenter := factory.Init()

	eAuth := e.Group("/auth")
	eAuth.POST("", presenter.AuthHandler.LoginHandler)
	eAuth.PUT("", presenter.AuthHandler.ReLoginHandler)
	eAuth.DELETE("", presenter.AuthHandler.LogoutHandler)

	eUsers := e.Group("/users")
	eUsers.POST("", presenter.UserHandler.RegisterUserHandler)
	eUsers.GET("", presenter.UserHandler.GetAllUsersHandler)
	eUsers.GET("/:userId", presenter.UserHandler.GetUserByIdHandler)
	eUsers.GET("/:userId/following", presenter.UserHandler.GetUserFollowingByIdHandler)
	eUsers.GET("/:userId/followers", presenter.UserHandler.GetUserFollowersByIdHandler)

	eArticles := e.Group("/articles")
	configJWT := middleware.JWTConfig{
		SigningKey: []byte(constants.ACCESS_TOKEN_KEY),
		Claims:     &middlewares.JwtCustomClaims{},
	}
	eArticles.Use(middleware.JWTWithConfig(configJWT))
	eArticles.POST("", presenter.ArticleHandler.CreateArticleHandler)
	eArticles.GET("", presenter.ArticleHandler.GetAllArticleHandler)
	eArticles.GET("/:articleId", presenter.ArticleHandler.GetArticleByIdHandler)
	eArticles.PATCH("/:articleId", presenter.ArticleHandler.GetArticleByIdHandler)
	eArticles.DELETE("/:articleId", presenter.ArticleHandler.GetArticleByIdHandler)
	// middlewares.Logger(e)
	return n

}
