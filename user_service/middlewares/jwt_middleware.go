package middlewares

import (
	"managerservice/config"
	"managerservice/exceptions"
	"strings"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func JWTMiddleware(c *fiber.Ctx) error {
	jwtware.New(jwtware.Config{
		ErrorHandler: exceptions.CustomErrorHandling,
		SigningKey:   []byte(config.GetConfig().JWT_SECRET),
	})
	return c.Next()
}

func validateToken(tokenString string) (bool, *jwt.Token) {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return []byte(config.GetConfig().JWT_SECRET), nil
	})

	if err != nil {
		panic(err)
	}

	return token.Valid, token
}

func ExtractToken(c *fiber.Ctx) (int, string) {
	tokenString := c.GetReqHeaders()["Authorization"]
	tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")
	isValid, token := validateToken(tokenString)
	if isValid {
		claims := token.Claims.(jwt.MapClaims)
		id := claims["id"].(float64)
		role := claims["role"].(string)
		return int(id), role
	}
	return 0, ""
}
