package middlewares

import (
	"managerservice/utils/helpers"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func IsAdmin(c *fiber.Ctx) error {
	id, role := ExtractToken(c)

	if role != "admin" || id < 1 {
		return c.Status(http.StatusMethodNotAllowed).JSON(helpers.FailedResponse("role not admin"))
	}

	return c.Next()
}
