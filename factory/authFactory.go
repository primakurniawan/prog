package factory

import (
	"prog/db"
	"prog/features/auth/business"
	"prog/features/auth/data"
	"prog/features/auth/presentation"
)

type AuthPresenter struct {
	AuthHandler presentation.AuthHandler
}

func InitAuth() AuthPresenter {
	authData := data.NewMysqlAuthRepository(db.DB)
	authBusiness := business.NewAuthBusiness(authData)
	authPresentation := presentation.NewAuthHandler(authBusiness)

	return AuthPresenter{
		AuthHandler: *authPresentation,
	}
}
