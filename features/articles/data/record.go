package data

import (
	"prog/features/articles"
	userData "prog/features/users/data"
	"time"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	ID        int
	Title     string
	Image     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int
	User      userData.User
	Tags      []Tag `gorm:"many2many:article_tags;"`
}

type Tag struct {
	gorm.Model
	ID    int
	Title string `gorm:"unique"`
}

func ToTagRecord(tag articles.TagCore) Tag {
	return Tag{
		ID:    tag.ID,
		Title: tag.Title,
	}
}

func ToArticleRecord(article articles.ArticleCore) Article {
	return Article{
		ID:        article.ID,
		Title:     article.Title,
		Image:     article.Image,
		Content:   article.Content,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
		UserID:    article.UserID,
		User:      userData.ToUserRecord(article.User),
		Tags:      ToTagsRecordList(article.Tags),
	}
}

func ToArticleCore(article Article) articles.ArticleCore {
	return articles.ArticleCore{
		ID:        article.ID,
		Title:     article.Title,
		Image:     article.Image,
		Content:   article.Content,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
		User:      userData.ToUserCore(article.User),
		Tags:      ToTagsCoreList(article.Tags),
	}
}

func ToTagCore(tag Tag) articles.TagCore {
	return articles.TagCore{
		ID:    tag.ID,
		Title: tag.Title,
	}
}

func ToArticleCoreList(aList []Article) []articles.ArticleCore {
	convertedArticle := []articles.ArticleCore{}

	for _, article := range aList {
		convertedArticle = append(convertedArticle, ToArticleCore(article))
	}

	return convertedArticle
}

func ToTagsCoreList(tList []Tag) []articles.TagCore {
	convertedTag := []articles.TagCore{}

	for _, tag := range tList {
		convertedTag = append(convertedTag, ToTagCore(tag))
	}

	return convertedTag
}

func ToTagsRecordList(tList []articles.TagCore) []Tag {
	convertedUser := []Tag{}

	for _, tag := range tList {
		convertedUser = append(convertedUser, ToTagRecord(tag))
	}

	return convertedUser
}
