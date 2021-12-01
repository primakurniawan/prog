package business

import (
	"prog/features/auth"
	"prog/features/users"
)

type authUsecase struct {
	AuthData auth.Data
}

func NewAuthBusiness(authData auth.Data) auth.Business {
	return &authUsecase{AuthData: authData}
}

func (au *authUsecase) AddRefreshToken(data auth.Core) error {
	err := au.AuthData.AddRefreshToken(data)
	if err != nil {
		return err
	}
	return nil
}

func (au *authUsecase) VerifyRefreshToken(data auth.Core) error {
	err := au.AuthData.VerifyRefreshToken(data)
	if err != nil {
		return err
	}
	return nil
}

func (au *authUsecase) DeleteRefreshToken(data auth.Core) error {

	err := au.AuthData.DeleteRefreshToken(data)
	if err != nil {
		return err
	}
	return nil

}

func (au *authUsecase) VerifyUserCredential(data users.Core) (userId int, err error) {

	userId, err = au.AuthData.VerifyUserCredential(data)
	if err != nil {
		return 0, err
	}
	return userId, nil

}
