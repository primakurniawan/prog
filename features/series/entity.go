package series

import "time"

type SeriesCore struct {
	ID          int
	Title       string
	Description string
	UserID      int
	User        UserCore
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ArticlesSeriesCore struct {
	ArticleId int
	Article   ArticleCore
	SeriesId  int
	Series    SeriesCore
}

type UserCore struct {
	ID       int
	Email    string
	Fullname string
	Image    string
}

type ArticleCore struct {
	ID        int
	Title     string
	Image     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserId    int
	User      UserCore
	Tags      []TagCore
}

type TagCore struct {
	ID    int
	Title string
}

type Business interface {
	CreateSeries(data SeriesCore) error
	DeleteSeries(seriesId int) error
	UpdateSeriesById(seriesId int, data SeriesCore) error
	AddArticleSeries(data ArticlesSeriesCore) error
	GetAllSeries() ([]SeriesCore, error)
	GetAllArticleSeries(seriesId int) ([]ArticleCore, error)
	VerifySeriesOwner(seriesId, userId int) error
}

type Data interface {
	CreateSeries(data SeriesCore) error
	DeleteSeries(seriesId int) error
	UpdateSeriesById(seriesId int, data SeriesCore) error
	AddArticleSeries(data ArticlesSeriesCore) error
	GetAllSeries() ([]SeriesCore, error)
	GetAllArticleSeries(seriesId int) ([]ArticleCore, error)
	VerifySeriesOwner(seriesId, userId int) error
}
