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
	Article   articles.Core
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	AddComment(content string, articleId, userId int) error
	GetArticleComments(articleId int) ([]Core, error)
	UpdateComment(content string, commentId, userId int) error
	DeleteComment(commentId, userId int) error
}

type Data interface {
	AddComment(content string, articleId, userId int) error
	GetArticleComments(articleId int) ([]Core, error)
	DeleteComment(commentId int) error
	UpdateComment(commentId int, content string) error
	VerifyCommentOwner(commentId, userId int) error
}
