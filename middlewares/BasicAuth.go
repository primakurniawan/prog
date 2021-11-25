package middlewares

import (
	"prog/db"
	"prog/models"

	"github.com/labstack/echo/v4"
)

func BasicAuth(email, password string, c echo.Context) (bool, error) {
	var user models.User
	err := db.DB.Where("email = ? AND password = ?", email, password).First(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
