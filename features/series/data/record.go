package data

import (
	"prog/features/articles"
	articleData "prog/features/articles/data"
	"prog/features/series"
	"prog/features/users"
	userData "prog/features/users/data"
	"time"

	"gorm.io/gorm"
)

type Series struct {
	gorm.Model
	ID          int
	Title       string
	Description string
	UserID      int
	User        userData.User
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ArticleSeries struct {
	gorm.Model
	ArticleId int
	Article   articleData.Article
	SeriesId  int
	Series    Series
}

func toArticlesSeriesRecord(articlesSeries series.ArticlesSeriesCore) ArticleSeries {
	return ArticleSeries{
		ArticleId: articlesSeries.ArticleId,
		SeriesId:  articlesSeries.SeriesId,
	}
}

func toSeriesRecord(series series.SeriesCore) Series {
	return Series{
		ID:          series.ID,
		Title:       series.Title,
		Description: series.Description,
		UserID:      series.UserID,
		User:        toUserRecord(series.User),
	}
}

func toSeriesCore(seriesRecord Series) series.SeriesCore {
	return series.SeriesCore{
		ID:          seriesRecord.ID,
		Title:       seriesRecord.Title,
		Description: seriesRecord.Description,
		UserID:      seriesRecord.UserID,
		User:        toUserCore(seriesRecord.User),
	}
}

func toUserRecord(user users.Core) userData.User {
	return userData.User{
		ID:       user.ID,
		Email:    user.Email,
		Fullname: user.Fullname,
		Image:    user.Email,
	}
}

func toUserCore(user userData.User) users.Core {
	return users.Core{
		ID:       user.ID,
		Email:    user.Email,
		Fullname: user.Fullname,
		Image:    user.Image,
	}
}

func toSeriesCoreList(sList []Series) []series.SeriesCore {
	convertedSeries := []series.SeriesCore{}

	for _, series := range sList {
		convertedSeries = append(convertedSeries, toSeriesCore(series))
	}

	return convertedSeries
}

func ToArticleCoreList(aList []ArticleSeries) []articles.Core {
	convertedArticle := []articles.Core{}

	for _, articleSeries := range aList {
		convertedArticle = append(convertedArticle, articleData.ToArticleCore(articleSeries.Article))
	}

	return convertedArticle
}
