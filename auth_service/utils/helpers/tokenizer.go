package helpers

import (
	"authentication/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(id int, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GetConfig().JWT_SECRET))
}
