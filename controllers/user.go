package controllers

import (
	"net/http"
	"prog/models"

	"github.com/labstack/echo"
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
