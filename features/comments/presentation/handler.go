package presentation

import (
	"net/http"
	"prog/features/articles"
	"prog/features/comments"
	"prog/features/comments/presentation/request"
	"prog/features/comments/presentation/response"
	"prog/middlewares"

	"strconv"

	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	CommentBusiness comments.Business
	ArticleBusiness articles.Business
}

func NewCommentHandler(commentBusiness comments.Business, articleBusiness articles.Business) *CommentHandler {
	return &CommentHandler{CommentBusiness: commentBusiness, ArticleBusiness: articleBusiness}
}

func (uh *CommentHandler) AddComment(e echo.Context) error {

	articleId, err := strconv.Atoi(e.Param("articleId"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not comment article",
			"err":     "articleId must be integer",
		})
	}

	userId, err := middlewares.VerifyAccessToken(e)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not comment article",
			"err":     err.Error(),
		})
	}

	comment := request.CommentRequest{}
	e.Bind(&comment)

	err = uh.CommentBusiness.AddComment(comment.ToCommentCore(articleId, userId))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not comment article",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "comment article",
	})
}

func (uh *CommentHandler) UpdateComment(e echo.Context) error {
	articleId, err := strconv.Atoi(e.Param("articleId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not delete comment",
			"err":     "articleId must be integer",
		})
	}

	_, err = uh.ArticleBusiness.GetArticleById(articleId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not delete comment",
			"err":     "articleId must be integer",
		})
	}

	commentId, err := strconv.Atoi(e.Param("commentId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not update comment",
			"err":     "commentId must be integer",
		})
	}

	userId, err := middlewares.VerifyAccessToken(e)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not update comment",
			"err":     err.Error(),
		})
	}

	err = uh.CommentBusiness.VerifyCommentOwner(commentId, userId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not update comment",
			"err":     err.Error(),
		})
	}

	comment := request.CommentRequest{}
	e.Bind(&comment)
	err = uh.CommentBusiness.UpdateComment(commentId, comment.ToCommentCore(articleId, userId))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not update comment",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "update comment",
	})

}

func (uh *CommentHandler) DeleteComment(e echo.Context) error {
	articleId, err := strconv.Atoi(e.Param("articleId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not delete comment",
			"err":     "articleId must be integer",
		})
	}

	_, err = uh.ArticleBusiness.GetArticleById(articleId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not delete comment",
			"err":     "articleId must be integer",
		})
	}

	commentId, err := strconv.Atoi(e.Param("commentId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not delete comment",
			"err":     "commentId must be integer",
		})
	}

	userId, err := middlewares.VerifyAccessToken(e)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not delete comment",
			"err":     err.Error(),
		})
	}

	err = uh.CommentBusiness.VerifyCommentOwner(commentId, userId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not delete comment",
			"err":     err.Error(),
		})
	}

	err = uh.CommentBusiness.DeleteComment(commentId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not delete comment",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "delete comment",
	})

}

func (alh *CommentHandler) GetArticleComments(e echo.Context) error {
	articleId, err := strconv.Atoi(e.Param("articleId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not get article comments",
			"err":     err.Error(),
		})
	}
	data, err := alh.CommentBusiness.GetArticleComments(articleId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not get article comments",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   response.ToCommentResponseList(data),
	})

}
