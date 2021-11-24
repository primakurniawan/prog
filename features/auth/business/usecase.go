package business

import (
	"prog/features/auth"
	"prog/features/users"
	"prog/middlewares"
)

type authUsecase struct {
	AuthData auth.Data
}

func NewAuthBusiness(authData auth.Data) auth.Business {
	return &authUsecase{AuthData: authData}
}

func (au *authUsecase) Login(data users.Core) (accessTokenCore auth.Core, refreshTokenCore auth.Core, err error) {
	user, err := au.AuthData.VerifyUserCredential(data)
	if err != nil {
		return auth.Core{}, auth.Core{}, err
	}
	accessToken, err := middlewares.CreateToken(user.ID)
	if err != nil {
		return auth.Core{}, auth.Core{}, err
	}
	accessTokenCore = auth.Core{
		Token: accessToken,
	}

	refreshToken, err := middlewares.CreateRefreshToken(user.ID)
	if err != nil {
		return auth.Core{}, auth.Core{}, err
	}
	err = au.AuthData.AddRefreshToken(auth.Core{
		Token: refreshToken,
	})
	if err != nil {
		return auth.Core{}, auth.Core{}, err
	}
	refreshTokenCore = auth.Core{
		Token: refreshToken,
	}
	return accessTokenCore, refreshTokenCore, nil
}

func (au *authUsecase) ReLogin(data auth.Core, userId int) (accessTokenCore auth.Core, err error) {
	err = au.AuthData.VerifyRefreshToken(data)
	if err != nil {
		return auth.Core{}, err
	}
	accessToken, err := middlewares.CreateToken(userId)
	if err != nil {
		return auth.Core{}, err
	}
	return auth.Core{
		Token: accessToken,
	}, nil
}

func (au *authUsecase) Logout(data auth.Core) (err error) {
	err = au.AuthData.VerifyRefreshToken(data)
	if err != nil {
		return err
	}
	err = au.AuthData.DeleteRefreshToken(data)
	if err != nil {
		return err
	}
	return nil
}
