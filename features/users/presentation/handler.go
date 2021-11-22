package presentation

import (
	"fmt"
	"net/http"
	"prog/features/users"
	"prog/features/users/presentation/request"
	"prog/features/users/presentation/response"
	"strconv"

	"github.com/labstack/echo"
)

type UserHandler struct {
	UserBusiness users.Business
}

func NewUserHandler(userBusiness users.Business) *UserHandler {
	return &UserHandler{UserBusiness: userBusiness}
}

func (uh *UserHandler) RegisterUserHandler(e echo.Context) error {
	userData := request.UserRequest{}

	err := e.Bind(&userData)
	fmt.Println("data in handler =====", userData)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	err = uh.UserBusiness.RegisterUser(userData.ToUserCore())
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}

func (uh *UserHandler) GetAllUserHandler(e echo.Context) error {
	fullname := e.QueryParam("fullname")
	data, err := uh.UserBusiness.GetUsersByFullname(fullname)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToUserResponseList(data),
	})

}

func (uh *UserHandler) GetUserByIdHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("userId"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	data, err := uh.UserBusiness.GetUserById(id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToUserResponse(data),
	})

}
