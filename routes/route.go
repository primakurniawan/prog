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

	configJWT := middleware.JWTConfig{
		SigningKey: []byte(constants.ACCESS_TOKEN_KEY),
		Claims:     &middlewares.JwtCustomClaims{},
	}

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
	eUsers.PUT("/:userId/follow", presenter.FollowHandler.FollowUser)
	eUsers.DELETE("/:userId/follow", presenter.FollowHandler.UnfollowUser)
	eUsers.GET("/:userId/likes", presenter.ArticleLikesHandler.GetLikedArticles)
	eUsers.GET("/:userId/articles", presenter.ArticleHandler.GetAllUserArticlesHandler)

	// eUser := e.Group("/user")
	// eUser.GET("", presenter.UserHandler.GetUserByIdHandler)

	eArticles := e.Group("/articles")
	eArticles.POST("", presenter.ArticleHandler.CreateArticleHandler, middleware.JWTWithConfig(configJWT))
	eArticles.GET("", presenter.ArticleHandler.GetAllArticleHandler)
	eArticles.GET("/:articleId", presenter.ArticleHandler.GetArticleByIdHandler)
	eArticles.GET("/:articleId/likes", presenter.ArticleLikesHandler.GetLikingUsers)
	eArticles.PUT("/:articleId/likes", presenter.ArticleLikesHandler.LikeArticle, middleware.JWTWithConfig(configJWT))
	eArticles.DELETE("/:articleId/likes", presenter.ArticleLikesHandler.UnlikeArticle, middleware.JWTWithConfig(configJWT))
	eArticles.PATCH("/:articleId", presenter.ArticleHandler.UpdateArticleByIdHandler, middleware.JWTWithConfig(configJWT))
	eArticles.DELETE("/:articleId", presenter.ArticleHandler.DeleteArticleByIdHandler, middleware.JWTWithConfig(configJWT))
	eArticles.GET("/:articleId/comments", presenter.CommentHandler.GetArticleComments)
	eArticles.POST("/:articleId/comments", presenter.CommentHandler.AddComment, middleware.JWTWithConfig(configJWT))
	eArticles.PATCH("/:articleId/comments/:commentId", presenter.CommentHandler.UpdateComment, middleware.JWTWithConfig(configJWT))
	eArticles.DELETE("/:articleId/comments/:commentId", presenter.CommentHandler.DeleteComment, middleware.JWTWithConfig(configJWT))
	// middlewares.Logger(e)
	return n

}
