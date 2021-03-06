package routes

import (
	"net/http"
	"os"
	"prog/factory"
	"prog/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var ACCESS_TOKEN_KEY string = os.Getenv("ACCESS_TOKEN_KEY")

func New() *echo.Echo {
	n := echo.New()
	e := n.Group("/v1")

	configJWT := middleware.JWTConfig{
		SigningKey: []byte(ACCESS_TOKEN_KEY),
		Claims:     &middlewares.JwtCustomClaims{},
	}

	presenter := factory.Init()

	n.GET("/api/check/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK!")
	})
	eAuth := e.Group("/auth")
	eAuth.POST("", presenter.AuthHandler.LoginHandler)
	eAuth.PUT("", presenter.AuthHandler.ReLoginHandler)
	eAuth.DELETE("", presenter.AuthHandler.LogoutHandler)

	eUsers := e.Group("/users")
	eUsers.POST("", presenter.UserHandler.RegisterUserHandler)
	eUsers.GET("", presenter.UserHandler.GetAllUsersHandler)
	eUsers.GET("/:userId", presenter.UserHandler.GetUserByIdHandler)
	eUsers.GET("/:userId/following", presenter.FollowHandler.GetFollowingUsers)
	eUsers.GET("/:userId/followers", presenter.FollowHandler.GetFollowersUsers)
	eUsers.PUT("/:userId/follow", presenter.FollowHandler.FollowUser)
	eUsers.DELETE("/:userId/follow", presenter.FollowHandler.UnfollowUser)
	eUsers.GET("/:userId/likes", presenter.ArticleLikesHandler.GetLikedArticles)
	eUsers.GET("/:userId/articles", presenter.ArticleHandler.GetAllUserArticlesHandler)

	eUser := e.Group("/user")
	eUser.Use(middleware.JWTWithConfig(configJWT))
	eUser.GET("", presenter.UserHandler.GetUserHandler)
	eUser.PATCH("", presenter.UserHandler.UpdateUserHandler)
	eUser.DELETE("", presenter.UserHandler.DeleteUserHandler)

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

	eSeries := e.Group("/series")
	eSeries.GET("", presenter.SeriesHandler.GetAllSeriesHandler)
	eSeries.POST("", presenter.SeriesHandler.CreateSeriesHandler, middleware.JWTWithConfig(configJWT))
	eSeries.GET("/:seriesId", presenter.SeriesHandler.GetSeriesByIdHandler)
	eSeries.DELETE("/:seriesId", presenter.SeriesHandler.DeleteSeriesHandler, middleware.JWTWithConfig(configJWT))
	eSeries.PATCH("/:seriesId", presenter.SeriesHandler.UpdateSeriesByIdHandler, middleware.JWTWithConfig(configJWT))
	eSeries.POST("/:seriesId", presenter.SeriesHandler.AddArticleSeriesHandler, middleware.JWTWithConfig(configJWT))
	eSeries.GET("/:seriesId/articles", presenter.SeriesHandler.GetAllArticlesSeriesHandler)

	e.GET("/news", presenter.NewsHandler.GetNewsHandler)

	return n

}
