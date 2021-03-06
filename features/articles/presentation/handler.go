package presentation

import (
	"net/http"
	"prog/features/articles"
	"prog/features/articles/presentation/request"
	"prog/features/articles/presentation/response"
	"prog/middlewares"

	"strconv"

	"github.com/labstack/echo/v4"
)

type ArticleHandler struct {
	ArticleBusiness articles.Business
}

func NewArticleHandler(articleBusiness articles.Business) *ArticleHandler {
	return &ArticleHandler{ArticleBusiness: articleBusiness}
}

func (uh *ArticleHandler) CreateArticleHandler(e echo.Context) error {
	articleData := request.ArticleRequest{}

	e.Bind(&articleData)

	userId, err := middlewares.VerifyAccessToken(e)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not create article",
			"err":     err.Error(),
		})
	}

	tags, err := uh.ArticleBusiness.CreateTags(articleData.ToTagCoreList())
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not create article",
			"err":     err.Error(),
		})
	}

	err = uh.ArticleBusiness.CreateArticle(articleData.ToArticleCore(tags, userId))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not create article",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "new article is created",
	})
}

func (ah *ArticleHandler) GetAllArticleHandler(e echo.Context) error {
	data, err := ah.ArticleBusiness.GetAllArticles()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not get all articles",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   response.ToArticleResponseList(data),
	})

}

func (ah *ArticleHandler) GetArticleByIdHandler(e echo.Context) error {
	articleId, err := strconv.Atoi(e.Param("articleId"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not get article",
			"err":     err.Error(),
		})
	}

	data, err := ah.ArticleBusiness.GetArticleById(articleId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not get article",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   response.ToArticleResponse(data),
	})

}

func (ah *ArticleHandler) UpdateArticleByIdHandler(e echo.Context) error {
	articleId, err := strconv.Atoi(e.Param("articleId"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not update article",
			"err":     err.Error(),
		})
	}

	articleData := request.ArticleRequest{}
	e.Bind(&articleData)

	tags, err := ah.ArticleBusiness.CreateTags(articleData.ToTagCoreList())
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not create article",
			"err":     err.Error(),
		})
	}

	userId, err := middlewares.VerifyAccessToken(e)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not create article",
			"err":     err.Error(),
		})
	}

	err = ah.ArticleBusiness.VerifyArticleOwner(articleId, userId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not create article",
			"err":     err.Error(),
		})
	}

	err = ah.ArticleBusiness.UpdateArticleById(articleId, articleData.ToArticleCore(tags, userId))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not update article",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "update article",
	})

}

func (uh *ArticleHandler) DeleteArticleByIdHandler(e echo.Context) error {
	articleId, err := strconv.Atoi(e.Param("articleId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not delete article",
			"err":     err.Error(),
		})
	}

	userId, err := middlewares.VerifyAccessToken(e)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not delete article",
			"err":     err.Error(),
		})
	}

	err = uh.ArticleBusiness.VerifyArticleOwner(articleId, userId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not create article",
			"err":     err.Error(),
		})
	}

	err = uh.ArticleBusiness.DeleteArticleById(articleId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not delete article",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "delete article",
	})

}

func (ah *ArticleHandler) GetAllUserArticlesHandler(e echo.Context) error {

	userId, err := strconv.Atoi(e.Param("userId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not get all user articles",
			"err":     err.Error(),
		})
	}
	data, err := ah.ArticleBusiness.GetAllUserArticles(userId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not get all user articles",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   response.ToArticleResponseList(data),
	})

}
