package request

import (
	"prog/features/auth"
	"prog/features/users"
)

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

func (requestData *UserRequest) ToUserCore() users.Core {
	return users.Core{
		Email:    requestData.Email,
		Password: requestData.Password,
	}
}

func (requestData *TokenRequest) ToTokenCore() auth.Core {
	return auth.Core{
		Token: requestData.RefreshToken,
	}
}
