package articles

import (
	"prog/features/users"
	"time"
)

type ArticleCore struct {
	ID        int
	Title     string
	Image     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int
	User      users.Core
	Tags      []TagCore
}

type TagCore struct {
	ID    int
	Title string
}

type Business interface {
	CreateTags(tags []TagCore) ([]TagCore, error)
	CreateArticle(data ArticleCore) error
	GetAllArticles() ([]ArticleCore, error)
	GetArticleById(articleId int) (ArticleCore, error)
	UpdateArticleById(articleId int, data ArticleCore) error
	DeleteArticleById(articleId int) error
	VerifyArticleOwner(articleId int, userId int) error
	GetAllUserArticles(userId int) ([]ArticleCore, error)
}

type Data interface {
	CreateTags(tags []TagCore) ([]TagCore, error)
	CreateArticle(data ArticleCore) error
	GetAllArticles() ([]ArticleCore, error)
	GetArticleById(articleId int) (ArticleCore, error)
	UpdateArticleById(articleId int, data ArticleCore) error
	DeleteArticleById(articleId int) error
	VerifyArticleOwner(articleId int, userId int) error
	GetAllUserArticles(userId int) ([]ArticleCore, error)
}
