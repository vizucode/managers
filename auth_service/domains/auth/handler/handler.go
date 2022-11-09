package userhandler

import (
	usercore "authentication/domains/auth/core"
	"authentication/exceptions"
	"authentication/utils/helpers"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	service usercore.IServiceAuth
}

var validate = validator.New()

func New(service usercore.IServiceAuth) *authHandler {
	return &authHandler{
		service: service,
	}
}

func (h *authHandler) Login(c *fiber.Ctx) error {
	request := RequestVerify{}

	err := c.BodyParser(&request)
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	err = validate.Struct(&request)
	if err != nil {
		return err
	}

	result := h.service.Verify(usercore.Core{
		Email:    request.Email,
		Password: request.Password,
	})

	if result.Id > 0 {
		return c.Status(http.StatusOK).JSON(helpers.SuccessGetResponseData(Response{
			Name:  result.Name,
			Token: result.Token,
		}))
	} else {
		return c.Status(http.StatusUnauthorized).JSON(helpers.SuccessActionResponse("unauthorized"))
	}
}
