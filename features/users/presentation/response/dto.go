package response

import (
	"prog/features/users"
)

type UserResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Image    string `json:"image"`
}

func ToUserResponse(user users.Core) UserResponse {
	return UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		Fullname: user.Fullname,
		Image:    user.Image,
	}
}

func ToUserResponseList(userList []users.Core) []UserResponse {
	convertedUser := []UserResponse{}
	for _, user := range userList {
		convertedUser = append(convertedUser, ToUserResponse(user))
	}

	return convertedUser
}
