package middleware

import (
	"laptop_catalog/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(adminID int, name string) (string, error) {

	claims := jwt.MapClaims{}

	claims["authorized"] = true
	claims["adminID"] = adminID
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constants.SECRET_JWT))
}
