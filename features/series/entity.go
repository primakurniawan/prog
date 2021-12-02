package series

import (
	"prog/features/articles"
	"prog/features/users"
	"time"
)

type SeriesCore struct {
	ID          int
	Title       string
	Description string
	UserID      int
	User        users.Core
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ArticlesSeriesCore struct {
	ArticleId int
	Article   articles.ArticleCore
	SeriesId  int
	Series    SeriesCore
}

type Business interface {
	CreateSeries(data SeriesCore) error
	DeleteSeries(seriesId int) error
	UpdateSeriesById(seriesId int, data SeriesCore) error
	AddArticleSeries(data ArticlesSeriesCore) error
	GetAllSeries() ([]SeriesCore, error)
	GetSeriesById(seriesId int) (SeriesCore, error)
	GetAllArticleSeries(seriesId int) ([]articles.ArticleCore, error)
	VerifySeriesOwner(seriesId, userId int) error
}

type Data interface {
	CreateSeries(data SeriesCore) error
	DeleteSeries(seriesId int) error
	UpdateSeriesById(seriesId int, data SeriesCore) error
	AddArticleSeries(data ArticlesSeriesCore) error
	GetAllSeries() ([]SeriesCore, error)
	GetSeriesById(seriesId int) (SeriesCore, error)
	GetAllArticleSeries(seriesId int) ([]articles.ArticleCore, error)
	VerifySeriesOwner(seriesId, userId int) error
}
