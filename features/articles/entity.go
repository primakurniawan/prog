package articles

import (
	"prog/features/users"
	"time"
)

type Core struct {
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
	CreateArticle(data Core, userId int) error
	GetAllArticles() ([]Core, error)
	GetArticleById(articleId int) (Core, error)
	UpdateArticleById(articleId int, data Core, userId int) error
	DeleteArticleById(articleId int, userId int) error
	GetAllUserArticles(userId int) ([]Core, error)
}

type Data interface {
	CreateTags(tags []TagCore) ([]TagCore, error)
	CreateArticle(data Core, userId int, tags []TagCore) error
	GetAllArticles() ([]Core, error)
	GetArticleById(articleId int) (Core, error)
	UpdateArticleById(articleId int, data Core) error
	DeleteArticleById(articleId int) error
	VerifyArticleOwner(articleId int, userId int) error
	GetAllUserArticles(userId int) ([]Core, error)
}
