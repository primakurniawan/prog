package data

import (
	"prog/features/articles"
	"prog/features/likes"
	"time"

	"gorm.io/gorm"
)

type ArticleLikes struct {
	gorm.Model
	ArticleID int
	Article   Article
	UserID    int
	User      User
}

type Article struct {
	gorm.Model
	ID        int
	Title     string
	Image     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int
	User      User
	Tags      []Tag `gorm:"many2many:article_tags;"`
}

type User struct {
	gorm.Model
	ID       int
	Email    string
	Fullname string
	Image    string
}

type Tag struct {
	gorm.Model
	ID    int
	Title string `gorm:"unique"`
}

func toArticleLikesRecord(data likes.Core) ArticleLikes {
	return ArticleLikes{
		UserID:    data.UserId,
		ArticleID: data.ArticleId,
	}
}

func toUserRecord(user articles.UserCore) User {
	return User{
		ID:       user.ID,
		Email:    user.Email,
		Fullname: user.Fullname,
		Image:    user.Image,
	}
}

func toTagRecord(tag articles.TagCore) Tag {
	return Tag{
		ID:    tag.ID,
		Title: tag.Title,
	}
}

func toArticleRecord(article articles.Core) Article {
	return Article{
		ID:        article.ID,
		Title:     article.Title,
		Image:     article.Image,
		Content:   article.Content,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
		UserID:    article.UserId,
		User:      toUserRecord(article.User),
		Tags:      toTagsRecordList(article.Tags),
	}
}

func toArticleCore(article Article) likes.ArticleCore {
	return likes.ArticleCore{
		ID:        article.ID,
		Title:     article.Title,
		Image:     article.Image,
		Content:   article.Content,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
		User:      toUserCore(article.User),
		Tags:      toTagsCoreList(article.Tags),
	}
}

func toUserCore(user User) likes.UserCore {
	return likes.UserCore{
		ID:       user.ID,
		Email:    user.Email,
		Fullname: user.Fullname,
		Image:    user.Image,
	}
}

func toTagCore(tag Tag) likes.TagCore {
	return likes.TagCore{
		ID:    tag.ID,
		Title: tag.Title,
	}
}

func toUserCoreList(aList []ArticleLikes) []likes.UserCore {
	convertedUser := []likes.UserCore{}

	for _, user := range aList {
		convertedUser = append(convertedUser, toUserCore(user.User))
	}

	return convertedUser
}
func toArticleCoreList(aList []ArticleLikes) []likes.ArticleCore {
	convertedArticle := []likes.ArticleCore{}

	for _, article := range aList {
		convertedArticle = append(convertedArticle, toArticleCore(article.Article))
	}

	return convertedArticle
}

func toTagsCoreList(tList []Tag) []likes.TagCore {
	convertedTag := []likes.TagCore{}

	for _, tag := range tList {
		convertedTag = append(convertedTag, toTagCore(tag))
	}

	return convertedTag
}

func toTagsRecordList(tList []articles.TagCore) []Tag {
	convertedUser := []Tag{}

	for _, tag := range tList {
		convertedUser = append(convertedUser, toTagRecord(tag))
	}

	return convertedUser
}
