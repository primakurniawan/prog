package response

import (
	"prog/features/follows"
)

type UserResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Image    string `json:"image"`
}

func toUserResponse(article follows.UserCore) UserResponse {
	return UserResponse{
		ID:       article.ID,
		Email:    article.Email,
		Fullname: article.Fullname,
		Image:    article.Image,
	}
}

func ToUserResponseList(userList []follows.UserCore) []UserResponse {
	convertedArticle := []UserResponse{}
	for _, user := range userList {
		convertedArticle = append(convertedArticle, toUserResponse(user))
	}

	return convertedArticle
}
