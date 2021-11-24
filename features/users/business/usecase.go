package business

import (
	"prog/features/users"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	UserData users.Data
}

func NewUserBusiness(userData users.Data) users.Business {
	return &userUsecase{UserData: userData}
}

func (uu *userUsecase) RegisterUser(data users.Core) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser := users.Core{
		Email:    data.Email,
		Password: string(hashedPassword),
		Fullname: data.Fullname,
		Image:    data.Image,
	}
	err = uu.UserData.CreateUser(newUser)

	if err != nil {
		return err
	}

	return nil
}

func (uu *userUsecase) GetAllUsers() ([]users.Core, error) {
	users, err := uu.UserData.GetAllUsers()
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

func (uu *userUsecase) GetUserFollowingById(userId int) ([]users.Core, error) {
	users, err := uu.UserData.GetUserFollowingById(userId)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uu *userUsecase) GetUserFollowersById(userId int) ([]users.Core, error) {
	users, err := uu.UserData.GetUserFollowersById(userId)
	if err != nil {
		return nil, err
	}

	return users, nil
}
