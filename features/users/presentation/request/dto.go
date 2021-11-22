package request

import "prog/features/users"

type UserRequest struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	Image    string `json:"image"`
}

func (requestData *UserRequest) ToUserCore() users.Core {
	return users.Core{
		ID:       requestData.ID,
		Email:    requestData.Email,
		Password: requestData.Password,
		Fullname: requestData.Fullname,
		Image:    requestData.Image,
	}
}
