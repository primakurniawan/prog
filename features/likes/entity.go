package likes

import (
	"prog/features/articles"
	"prog/features/users"
)

type Core struct {
	UserId    int
	User      users.Core
	ArticleId int
	Article   articles.Core
}

type Business interface {
	LikeArticle(articleId, userId int) error
	GetLikedArticles(userId int) ([]articles.Core, error)
	GetLikingUsers(articleId int) ([]users.Core, error)
	UnlikeArticle(articleId, userId int) error
}

type Data interface {
	LikeArticle(articleId, userId int) error
	GetLikedArticles(userId int) ([]articles.Core, error)
	GetLikingUsers(articleId int) ([]users.Core, error)
	UnlikeArticle(articleId, userId int) error
}
