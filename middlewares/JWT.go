package middlewares

import (
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var ACCESS_TOKEN_KEY string = os.Getenv("ACCESS_TOKEN_KEY")
var REFRESH_TOKEN_KEY string = os.Getenv("REFRESH_TOKEN_KEY")

type JwtCustomClaims struct {
	UserId int `json:"userId"`
	jwt.StandardClaims
}

func CreateToken(userId int) (string, error) {
	claims := &JwtCustomClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(ACCESS_TOKEN_KEY))
}

func CreateRefreshToken(userId int) (string, error) {
	claims := &JwtCustomClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(REFRESH_TOKEN_KEY))
}

func VerifyRefreshToken(refreshToken string) (userId int, err error) {
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		return []byte(REFRESH_TOKEN_KEY), nil
	}
	jwtToken, err := jwt.ParseWithClaims(refreshToken, &JwtCustomClaims{}, keyFunc)
	if err != nil {
		return 0, err
	}

	claims := jwtToken.Claims.(*JwtCustomClaims)
	userId = claims.UserId

	return userId, nil

}

func VerifyAccessToken(c echo.Context) (userId int, err error) {
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		return []byte(ACCESS_TOKEN_KEY), nil
	}
	accessToken := strings.Split(c.Request().Header.Get("Authorization"), " ")[1]
	jwtToken, err := jwt.ParseWithClaims(accessToken, &JwtCustomClaims{}, keyFunc)
	if err != nil {
		return 0, err
	}

	claims := jwtToken.Claims.(*JwtCustomClaims)
	userId = claims.UserId

	return userId, nil

}
