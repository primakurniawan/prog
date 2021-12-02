package data

import (
	articleData "prog/features/articles/data"
	"prog/features/comments"
	userData "prog/features/users/data"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID        int
	Content   string
	UserID    int
	User      userData.User
	ArticleID int
	Article   articleData.Article
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ToCommentsRecord(data comments.Core) Comment {
	return Comment{
		Content:   data.Content,
		UserID:    data.UserID,
		ArticleID: data.ArticleID,
	}
}

func ToCommentCore(comment Comment) comments.Core {
	return comments.Core{
		ID:        comment.ID,
		Content:   comment.Content,
		User:      userData.ToUserCore(comment.User),
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}
}

func ToCommentCoreList(cList []Comment) []comments.Core {
	convertedComment := []comments.Core{}

	for _, comment := range cList {
		convertedComment = append(convertedComment, ToCommentCore(comment))
	}

	return convertedComment
}
