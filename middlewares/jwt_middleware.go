package middlewares

import (
	"managers/config"
	"managers/exceptions"
	"time"

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

func CreateToken(id int, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GetConfig().JWT_SECRET))
}

func ExtractToken(c *fiber.Ctx) (int, string) {
	user := c.Locals("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		id := claims["id"].(float64)
		role := claims["role"].(string)
		return int(id), role
	}
	return 0, ""
}
