package factory

import (
	"prog/db"
	"prog/features/users/business"
	"prog/features/users/data"
	"prog/features/users/presentation"
)

type UserPresenter struct {
	UserHandler presentation.UserHandler
}

func InitUser() UserPresenter {
	userData := data.NewMysqlUserRepository(db.DB)
	userBusiness := business.NewUserBusiness(userData)
	userPresentation := presentation.NewUserHandler(userBusiness)

	return UserPresenter{
		UserHandler: *userPresentation,
	}
}
