package data

import (
	"errors"
	"prog/features/auth"
	"prog/features/users"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type mysqlAuthRepository struct {
	Conn *gorm.DB
}

func NewMysqlAuthRepository(conn *gorm.DB) auth.Data {
	return &mysqlAuthRepository{
		Conn: conn,
	}
}

func (ur *mysqlAuthRepository) AddRefreshToken(data auth.Core) error {
	recordData := toAuthRecord(data)
	err := ur.Conn.Create(&recordData).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *mysqlAuthRepository) VerifyRefreshToken(data auth.Core) error {
	recordData := toAuthRecord(data)
	err := ur.Conn.Where("token = ?", data.Token).First(&recordData).Error
	if err != nil {
		return err
	}
	if recordData.Token == "" {
		return errors.New("refresh token is not valid")
	}
	return nil
}

func (ur *mysqlAuthRepository) DeleteRefreshToken(data auth.Core) error {
	recordData := toAuthRecord(data)
	err := ur.Conn.Where("token = ?", data.Token).Delete(&recordData).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *mysqlAuthRepository) VerifyUserCredential(data users.Core) (users.Core, error) {

	recordData := toUserRecord(data)
	ur.Conn.Where("email = ?", data.Email).First(&recordData)
	if recordData.Password == "" && recordData.ID == 0 {
		return users.Core{
			ID: 0,
		}, errors.New("email is not registered yet")
	}
	if bcrypt.CompareHashAndPassword([]byte(recordData.Password), []byte(data.Password)) != nil {
		return users.Core{
			ID: 0,
		}, errors.New("password is not correct")
	}

	return toUserCore(recordData), nil
}
