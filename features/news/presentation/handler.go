package presentation

import (
	"net/http"
	"prog/features/news"

	"github.com/labstack/echo/v4"
)

type NewsHandler struct {
	newsService news.Business
}

func NewNewsHandler(ns news.Business) *NewsHandler {
	return &NewsHandler{ns}
}

func (ns *NewsHandler) GetNewsHandler(e echo.Context) error {
	q := e.QueryParam("q")
	data, err := ns.newsService.GetNews(q)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not get news",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   data,
	})
}
