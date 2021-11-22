package business

import (
	"prog/features/users"
)

type userUsecase struct {
	UserData users.Data
}

func NewUserBusiness(userData users.Data) users.Business {
	return &userUsecase{UserData: userData}
}

// GetUserById(id int) (Core, error)

func (uu *userUsecase) RegisterUser(data users.Core) error {
	err := uu.UserData.CreateUser(data)

	if err != nil {
		return err
	}

	return nil
}

func (uu *userUsecase) GetUsersByFullname(fullname string) ([]users.Core, error) {
	users, err := uu.UserData.GetUsersByFullname(fullname)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (us *userUsecase) GetUserById(id int) (users.Core, error) {
	userData, err := us.UserData.GetUserById(id)

	if err != nil {
		return users.Core{}, err
	}

	return userData, nil
}
