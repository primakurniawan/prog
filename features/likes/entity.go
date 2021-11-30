package likes

import "time"

type Core struct {
	UserId    int
	User      UserCore
	ArticleId int
	Article   ArticleCore
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
	LikeArticle(articleId, userId int) error
	GetLikedArticles(userId int) ([]ArticleCore, error)
	GetLikingUsers(articleId int) ([]UserCore, error)
	UnlikeArticle(articleId, userId int) error
}

type Data interface {
	LikeArticle(articleId, userId int) error
	GetLikedArticles(userId int) ([]ArticleCore, error)
	GetLikingUsers(articleId int) ([]UserCore, error)
	UnlikeArticle(articleId, userId int) error
}
