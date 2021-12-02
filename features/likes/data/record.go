package data

import (
	"prog/features/articles"
	articleData "prog/features/articles/data"
	"prog/features/likes"
	"prog/features/users"
	userData "prog/features/users/data"

	"gorm.io/gorm"
)

type ArticleLikes struct {
	gorm.Model
	ArticleID int
	Article   articleData.Article
	UserID    int
	User      userData.User
}

func toArticleLikesRecord(data likes.Core) ArticleLikes {
	return ArticleLikes{
		UserID:    data.UserId,
		ArticleID: data.ArticleId,
	}
}

func ToUserCoreList(aList []ArticleLikes) []users.Core {
	convertedUser := []users.Core{}

	for _, user := range aList {
		convertedUser = append(convertedUser, userData.ToUserCore(user.User))
	}

	return convertedUser
}
func ToArticleCoreList(aList []ArticleLikes) []articles.Core {
	convertedArticle := []articles.Core{}

	for _, article := range aList {
		convertedArticle = append(convertedArticle, articleData.ToArticleCore(article.Article))
	}

	return convertedArticle
}
