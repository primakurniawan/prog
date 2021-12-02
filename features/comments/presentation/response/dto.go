package response

import (
	"prog/features/comments"
	userResponse "prog/features/users/presentation/response"
	"time"
)

type CommentResponse struct {
	ID        int                       `json:"id"`
	Content   string                    `json:"content"`
	CreatedAt time.Time                 `json:"created_at"`
	UpdatedAt time.Time                 `json:"updated_at"`
	User      userResponse.UserResponse `json:"user"`
}

func ToCommentResponse(comment comments.Core) CommentResponse {
	return CommentResponse{
		ID:        comment.ID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
		User:      userResponse.ToUserResponse(comment.User),
	}
}

func ToCommentResponseList(commentList []comments.Core) []CommentResponse {
	convertedComment := []CommentResponse{}
	for _, comment := range commentList {
		convertedComment = append(convertedComment, ToCommentResponse(comment))
	}

	return convertedComment
}
