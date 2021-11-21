package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID       int
	Email    string `json:"email" form:"email"`
	Fullname string `json:"fullname" form:"fullname"`
}

var db *gorm.DB

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
	e := echo.New()

	e.GET("/users", getAllUsers)
	e.POST("/users", createUser)

	e.Start(":8000")
}

func getAllUsers(c echo.Context) error {
	users := []User{}

	db.Find(&users)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func createUser(c echo.Context) error {

	user := User{}
	c.Bind(&user)
	db.Create(&user)

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"status": "success",
		"user":   user,
	})

}
