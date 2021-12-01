package data

import (
	"prog/features/series"
	"time"

	"gorm.io/gorm"
)

type Series struct {
	gorm.Model
	ID          int
	Title       string
	Description string
	UserID      int
	User        User
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ArticleSeries struct {
	gorm.Model
	ArticleId int
	Article   Article
	SeriesId  int
	Series    Series
}

type User struct {
	gorm.Model
	ID       int
	Email    string
	Fullname string
	Image    string
}

type Article struct {
	gorm.Model
	ID        int
	Title     string
	Image     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserId    int
	User      User
	Tags      []Tag `gorm:"many2many:article_tags;"`
}

type Tag struct {
	gorm.Model
	ID    int
	Title string
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

func toUserRecord(user series.UserCore) User {
	return User{
		ID:       user.ID,
		Email:    user.Email,
		Fullname: user.Fullname,
		Image:    user.Email,
	}
}

func toUserCore(user User) series.UserCore {
	return series.UserCore{
		ID:       user.ID,
		Email:    user.Email,
		Fullname: user.Fullname,
		Image:    user.Image,
	}
}

func toArticleCore(article Article) series.ArticleCore {
	return series.ArticleCore{
		ID:      article.ID,
		Title:   article.Title,
		Image:   article.Image,
		Content: article.Content,
		User:    toUserCore(article.User),
		Tags:    toTagsCoreList(article.Tags),
	}
}

func toTagCore(tag Tag) series.TagCore {
	return series.TagCore{
		ID:    tag.ID,
		Title: tag.Title,
	}
}

func toTagsCoreList(tList []Tag) []series.TagCore {
	convertedTag := []series.TagCore{}

	for _, tag := range tList {
		convertedTag = append(convertedTag, toTagCore(tag))
	}

	return convertedTag
}

func toSeriesCoreList(sList []Series) []series.SeriesCore {
	convertedSeries := []series.SeriesCore{}

	for _, series := range sList {
		convertedSeries = append(convertedSeries, toSeriesCore(series))
	}

	return convertedSeries
}

func toArticleCoreList(aList []ArticleSeries) []series.ArticleCore {
	convertedArticle := []series.ArticleCore{}

	for _, articleSeries := range aList {
		convertedArticle = append(convertedArticle, toArticleCore(articleSeries.Article))
	}

	return convertedArticle
}
