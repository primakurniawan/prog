package controllers

import (
	"net/http"
	"prog/middlewares"
	"prog/models"

	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) error {
	users := models.FindAllUsers()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func CreateUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	user = models.CreateUser(user)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "success",
		"user":   user,
	})

}

func Login(c echo.Context) error {
	userId := models.Login(c.FormValue("email"), c.FormValue("password"))

	token, err := middlewares.CreateToken(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": "fail",
			"error":  err,
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "success",
		"token":  token,
	})
}
