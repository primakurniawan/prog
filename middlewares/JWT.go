package middlewares

import (
	"prog/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtCustomClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

func CreateToken(userId int) (string, error) {
	claims := &jwtCustomClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}
