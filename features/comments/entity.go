package comments

import (
	"prog/features/articles"
	"prog/features/users"
	"time"
)

type Core struct {
	ID        int
	Content   string
	UserID    int
	User      users.Core
	ArticleID int
	Article   articles.ArticleCore
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	AddComment(data Core) error
	GetArticleComments(articleId int) ([]Core, error)
	DeleteComment(commentId int) error
	UpdateComment(commentId int, data Core) error
	VerifyCommentOwner(commentId, userId int) error
}

type Data interface {
	AddComment(data Core) error
	GetArticleComments(articleId int) ([]Core, error)
	DeleteComment(commentId int) error
	UpdateComment(commentId int, data Core) error
	VerifyCommentOwner(commentId, userId int) error
}
