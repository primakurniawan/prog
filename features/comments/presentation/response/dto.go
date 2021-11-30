package response

import (
	"prog/features/comments"
	"time"
)

type CommentResponse struct {
	ID        int          `json:"id"`
	Content   string       `json:"content"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	User      UserResponse `json:"user"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Image    string `json:"image"`
}

func ToCommentResponse(comment comments.Core) CommentResponse {
	return CommentResponse{
		ID:        comment.ID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
		User:      toUserResponse(comment.User),
	}
}

func toUserResponse(user comments.UserCore) UserResponse {
	return UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		Fullname: user.Fullname,
		Image:    user.Image,
	}
}

func ToCommentResponseList(commentList []comments.Core) []CommentResponse {
	convertedComment := []CommentResponse{}
	for _, comment := range commentList {
		convertedComment = append(convertedComment, ToCommentResponse(comment))
	}

	return convertedComment
}
