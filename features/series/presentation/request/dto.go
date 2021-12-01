package request

import (
	"prog/features/series"
)

type SeriesRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ArticleSeriesRequest struct {
	ArticleId int `json:"article_id"`
}

func (requestData *SeriesRequest) ToSeriesCore(userId int) series.SeriesCore {
	return series.SeriesCore{
		Title:       requestData.Title,
		Description: requestData.Description,
		UserID:      userId,
	}
}

func (requestData *ArticleSeriesRequest) ToArticleSeriesCore(SeriesId int) series.ArticlesSeriesCore {
	return series.ArticlesSeriesCore{
		SeriesId:  SeriesId,
		ArticleId: requestData.ArticleId,
	}
}
