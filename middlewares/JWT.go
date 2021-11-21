package middlewares

import (
	"prog/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Minute).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}
