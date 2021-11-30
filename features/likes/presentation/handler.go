package presentation

import (
	"net/http"
	"prog/features/likes"
	"prog/features/likes/presentation/response"
	"prog/middlewares"

	"strconv"

	"github.com/labstack/echo/v4"
)

type ArticleLikesHandler struct {
	ArticleLikesBusiness likes.Business
}

func NewArticleLikesHandler(articleLikesBusiness likes.Business) *ArticleLikesHandler {
	return &ArticleLikesHandler{ArticleLikesBusiness: articleLikesBusiness}
}

func (uh *ArticleLikesHandler) LikeArticle(e echo.Context) error {

	articleId, err := strconv.Atoi(e.Param("articleId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not like article",
			"err":     "articleId must be integer",
		})
	}
	userId, err := middlewares.VerifyAccessToken(e)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not like article",
			"err":     err.Error(),
		})
	}
	err = uh.ArticleLikesBusiness.LikeArticle(articleId, userId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not like article",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "like article",
	})
}

func (uh *ArticleLikesHandler) UnlikeArticle(e echo.Context) error {
	articleId, err := strconv.Atoi(e.Param("articleId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not unlike article",
			"err":     "articleId must be integer",
		})
	}
	userId, err := middlewares.VerifyAccessToken(e)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not unlike article",
			"err":     err.Error(),
		})
	}
	err = uh.ArticleLikesBusiness.UnlikeArticle(articleId, userId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not unlike article",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "unlike article",
	})

}

func (alh *ArticleLikesHandler) GetLikedArticles(e echo.Context) error {
	userId, err := strconv.Atoi(e.Param("userId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not get liked articles",
			"err":     err.Error(),
		})
	}
	data, err := alh.ArticleLikesBusiness.GetLikedArticles(userId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not get liked articles",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   response.ToArticleResponseList(data),
	})

}

func (alh *ArticleLikesHandler) GetLikingUsers(e echo.Context) error {
	articleId, err := strconv.Atoi(e.Param("articleId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not get liking users",
			"err":     err.Error(),
		})
	}
	data, err := alh.ArticleLikesBusiness.GetLikingUsers(articleId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not get liking users",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   response.ToUserResponseList(data),
	})

}
