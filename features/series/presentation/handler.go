package presentation

import (
	"net/http"
	"prog/features/articles"
	articleResponse "prog/features/articles/presentation/response"
	"prog/features/series"
	"prog/features/series/presentation/request"
	"prog/features/series/presentation/response"
	"prog/middlewares"
	"strconv"

	"github.com/labstack/echo/v4"
)

type SeriesHandler struct {
	SeriesBusiness  series.Business
	ArticleBusiness articles.Business
}

func NewSeriesHandler(seriesBusiness series.Business, articleBusiness articles.Business) *SeriesHandler {
	return &SeriesHandler{SeriesBusiness: seriesBusiness, ArticleBusiness: articleBusiness}
}

func (sh *SeriesHandler) CreateSeriesHandler(e echo.Context) error {
	seriesData := request.SeriesRequest{}

	err := e.Bind(&seriesData)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not create series",
			"err":     err.Error(),
		})
	}

	userId, err := middlewares.VerifyAccessToken(e)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  "fail",
			"message": "can not create series",
			"err":     err.Error(),
		})
	}

	err = sh.SeriesBusiness.CreateSeries(seriesData.ToSeriesCore(userId))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not create series",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "success",
		"message": "new series is created",
	})

}

func (sh *SeriesHandler) DeleteSeriesHandler(e echo.Context) error {

	userId, err := middlewares.VerifyAccessToken(e)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  "fail",
			"message": "can not delete series",
			"err":     err.Error(),
		})
	}

	seriesId, err := strconv.Atoi(e.Param("seriesId"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not delete series",
			"err":     err.Error(),
		})
	}

	err = sh.SeriesBusiness.VerifySeriesOwner(seriesId, userId)
	if err != nil {
		return e.JSON(http.StatusForbidden, map[string]interface{}{
			"status":  "fail",
			"message": "can not delete series",
			"err":     err.Error(),
		})
	}

	err = sh.SeriesBusiness.DeleteSeries(seriesId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not delete series",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusNoContent, map[string]interface{}{
		"status":  "success",
		"message": "delete series",
	})

}

func (sh *SeriesHandler) UpdateSeriesByIdHandler(e echo.Context) error {
	userId, err := middlewares.VerifyAccessToken(e)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  "fail",
			"message": "can not update series",
			"err":     err.Error(),
		})
	}

	seriesId, err := strconv.Atoi(e.Param("seriesId"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not update series",
			"err":     err.Error(),
		})
	}

	err = sh.SeriesBusiness.VerifySeriesOwner(seriesId, userId)
	if err != nil {
		return e.JSON(http.StatusForbidden, map[string]interface{}{
			"status":  "fail",
			"message": "can not update series",
			"err":     err.Error(),
		})
	}

	seriesData := request.SeriesRequest{}
	e.Bind(&seriesData)

	err = sh.SeriesBusiness.UpdateSeriesById(seriesId, seriesData.ToSeriesCore(userId))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not update series",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusNoContent, map[string]interface{}{
		"status":  "success",
		"message": "update series",
	})

}

func (sh *SeriesHandler) GetAllSeriesHandler(e echo.Context) error {
	series, err := sh.SeriesBusiness.GetAllSeries()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not update series",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   response.ToSeriesResponseList(series),
	})

}

func (sh *SeriesHandler) GetSeriesByIdHandler(e echo.Context) error {
	seriesId, err := strconv.Atoi(e.Param("seriesId"))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not get series",
			"err":     err.Error(),
		})
	}
	series, err := sh.SeriesBusiness.GetSeriesById(seriesId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not get series",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   response.ToSeriesResponse(series),
	})

}

func (sh *SeriesHandler) GetAllArticlesSeriesHandler(e echo.Context) error {
	seriesId, err := strconv.Atoi(e.Param("seriesId"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not get articles series",
			"err":     err.Error(),
		})
	}

	articles, err := sh.SeriesBusiness.GetAllArticleSeries(seriesId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not get articles series",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   articleResponse.ToArticleResponseList(articles),
	})

}

func (sh *SeriesHandler) AddArticleSeriesHandler(e echo.Context) error {
	requestData := request.ArticleSeriesRequest{}

	userId, err := middlewares.VerifyAccessToken(e)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  "fail",
			"message": "can not create series",
			"err":     err.Error(),
		})
	}

	seriesId, err := strconv.Atoi(e.Param("seriesId"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not add article to series",
			"err":     err.Error(),
		})
	}

	e.Bind(&requestData)

	_, err = sh.ArticleBusiness.GetArticleById(requestData.ArticleId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not add article to series",
			"err":     err.Error(),
		})
	}

	err = sh.ArticleBusiness.VerifyArticleOwner(requestData.ArticleId, userId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not add article to series",
			"err":     err.Error(),
		})
	}

	err = sh.SeriesBusiness.VerifySeriesOwner(seriesId, userId)
	if err != nil {
		return e.JSON(http.StatusForbidden, map[string]interface{}{
			"status":  "fail",
			"message": "can not delete series",
			"err":     err.Error(),
		})
	}

	err = sh.SeriesBusiness.AddArticleSeries(requestData.ToArticleSeriesCore(seriesId))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not add article to series",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "add article to series",
	})

}

func (sh *SeriesHandler) DeleteArticleSeriesHandler(e echo.Context) error {
	requestData := request.ArticleSeriesRequest{}

	userId, err := middlewares.VerifyAccessToken(e)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  "fail",
			"message": "can not create series",
			"err":     err.Error(),
		})
	}

	seriesId, err := strconv.Atoi(e.Param("seriesId"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "can not add article to series",
			"err":     err.Error(),
		})
	}

	e.Bind(&requestData)

	_, err = sh.ArticleBusiness.GetArticleById(requestData.ArticleId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not add article to series",
			"err":     err.Error(),
		})
	}

	err = sh.ArticleBusiness.VerifyArticleOwner(requestData.ArticleId, userId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not add article to series",
			"err":     err.Error(),
		})
	}

	err = sh.SeriesBusiness.VerifySeriesOwner(seriesId, userId)
	if err != nil {
		return e.JSON(http.StatusForbidden, map[string]interface{}{
			"status":  "fail",
			"message": "can not delete series",
			"err":     err.Error(),
		})
	}

	err = sh.SeriesBusiness.DeleteArticleSeries(requestData.ToArticleSeriesCore(seriesId))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": "can not add article to series",
			"err":     err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "add article to series",
	})

}
