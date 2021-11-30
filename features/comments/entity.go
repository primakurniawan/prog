package comments

import "time"

type Core struct {
	ID        int
	Content   string
	UserID    int
	User      UserCore
	ArticleID int
	Article   ArticleCore
	CreatedAt time.Time
	UpdatedAt time.Time
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
