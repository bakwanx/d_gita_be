package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id int, fullname string) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["fullname"] = fullname
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	return token.SignedString("SECRET_JWT_D_GITA_H3H3HE")
}
