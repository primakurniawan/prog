package request

import "prog/features/comments"

type CommentRequest struct {
	Content string `json:"content"`
}

func (requestData *CommentRequest) ToCommentCore(articleId, userId int) comments.Core {
	return comments.Core{
		Content:   requestData.Content,
		ArticleID: articleId,
		UserID:    userId,
	}
}
