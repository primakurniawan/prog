package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Email    string `json:"email" form:"email"`
	Fullname string `json:"fullname" form:"fullname"`
}

func main() {
	e := echo.New()

	e.POST("/users", getUserById)

	e.Start(":8000")
}

func getUserById(c echo.Context) error {
	user := User{}

	c.Bind(&user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}
