package presentation

import (
	"net/http"
	"prog/features/auth"
	"prog/features/auth/presentation/request"
	"prog/features/auth/presentation/response"
	"prog/middlewares"

	"github.com/labstack/echo"
)

type AuthHandler struct {
	AuthBusiness auth.Business
}

func NewAuthHandler(authBusiness auth.Business) *AuthHandler {
	return &AuthHandler{AuthBusiness: authBusiness}
}

func (ah *AuthHandler) LoginHandler(e echo.Context) error {
	user := request.UserRequest{}
	e.Bind(&user)
	accessToken, refreshToken, err := ah.AuthBusiness.Login(user.ToUserCore())
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "cannot create acccess token and refresh token",
			"err":     err.Error(),
		})
	}
	return e.JSON(http.StatusBadRequest, map[string]interface{}{
		"status":  "success",
		"message": "create acccess token and refresh token",
		"data": response.AuthResponse{
			AccessToken:  accessToken.Token,
			RefreshToken: refreshToken.Token,
		},
	})
}

func (ah *AuthHandler) ReLoginHandler(e echo.Context) error {
	auth := request.TokenRequest{}
	e.Bind(&auth)
	userId, err := middlewares.VerifyRefreshToken(auth.RefreshToken)

	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "cannot create acccess token",
			"err":     err.Error(),
		})
	}

	accessToken, err := ah.AuthBusiness.ReLogin(auth.ToTokenCore(), userId)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "cannot create acccess token",
			"err":     err.Error(),
		})
	}
	return e.JSON(http.StatusBadRequest, map[string]interface{}{
		"status":  "success",
		"message": "create acccess token",
		"data": response.AuthRefreshResponse{
			AccessToken: accessToken.Token,
		},
	})

}

func (ah *AuthHandler) LogoutHandler(e echo.Context) error {
	auth := request.TokenRequest{}
	e.Bind(&auth)
	err := ah.AuthBusiness.Logout(auth.ToTokenCore())
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": "cannot delete refresh token",
			"err":     err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "delete refresh token",
	})

}
