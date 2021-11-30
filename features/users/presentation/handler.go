package presentation

import (
	"net/http"
	"prog/features/users"
	"prog/features/users/presentation/request"
	"prog/features/users/presentation/response"
	"strconv"

	"github.com/labstack/echo/v4"
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
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	userId, err := uh.UserBusiness.RegisterUser(userData.ToUserCore())
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	location := "http:/localhost:8000/v1/users/" + strconv.Itoa(userId)
	e.Response().Header().Set(echo.HeaderLocation, location)

	return e.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "success",
		"message": "new user is created",
	})
}

func (uh *UserHandler) GetAllUsersHandler(e echo.Context) error {
	data, err := uh.UserBusiness.GetAllUsers()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   response.ToUserResponseList(data),
	})

}

func (uh *UserHandler) GetUserByIdHandler(e echo.Context) error {
	userId, err := strconv.Atoi(e.Param("userId"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	data, err := uh.UserBusiness.GetUserById(userId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   response.ToUserResponse(data),
	})

}

func (uh *UserHandler) GetUserFollowingByIdHandler(e echo.Context) error {
	userId, _ := strconv.Atoi(e.Param("userId"))
	data, err := uh.UserBusiness.GetUserFollowingById(userId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   response.ToUserResponseList(data),
	})

}

func (uh *UserHandler) GetUserFollowersByIdHandler(e echo.Context) error {
	userId, _ := strconv.Atoi(e.Param("userId"))
	data, err := uh.UserBusiness.GetUserFollowersById(userId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   response.ToUserResponseList(data),
	})

}
