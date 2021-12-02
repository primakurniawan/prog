package data

import (
	"prog/features/articles"
	articleData "prog/features/articles/data"
	"prog/features/series"
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
	ArticleID int
	Article   articleData.Article
	SeriesID  int
	Series    Series
}

func toArticlesSeriesRecord(articlesSeries series.ArticlesSeriesCore) ArticleSeries {
	return ArticleSeries{
		ArticleID: articlesSeries.ArticleId,
		SeriesID:  articlesSeries.SeriesId,
	}
}

func toSeriesRecord(series series.SeriesCore) Series {
	return Series{
		ID:          series.ID,
		Title:       series.Title,
		Description: series.Description,
		UserID:      series.UserID,
		User:        userData.ToUserRecord(series.User),
	}
}

func toSeriesCore(seriesRecord Series) series.SeriesCore {
	return series.SeriesCore{
		ID:          seriesRecord.ID,
		Title:       seriesRecord.Title,
		Description: seriesRecord.Description,
		UserID:      seriesRecord.UserID,
		User:        userData.ToUserCore(seriesRecord.User),
	}
}

func toSeriesCoreList(sList []Series) []series.SeriesCore {
	convertedSeries := []series.SeriesCore{}

	for _, series := range sList {
		convertedSeries = append(convertedSeries, toSeriesCore(series))
	}

	return convertedSeries
}

func ToArticleCoreList(aList []ArticleSeries) []articles.ArticleCore {
	convertedArticle := []articles.ArticleCore{}

	for _, articleSeries := range aList {
		convertedArticle = append(convertedArticle, articleData.ToArticleCore(articleSeries.Article))
	}

	return convertedArticle
}
