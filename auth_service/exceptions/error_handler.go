package exceptions

import (
	"authentication/utils/helpers"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

/*
	Fungsi untuk membuat custom error.
*/
func CustomErrorHandling(ctx *fiber.Ctx, err error) error {
	if notFoundError(err, ctx) {
		return nil
	} else if badRequestError(err, ctx) {
		return nil

	} else if forbiddenError(err, ctx) {
		return nil

	} else if validationError(err, ctx) {
		return nil
	} else if jwtValidateError(err, ctx) {
		return nil
	} else {
		internalServerError(err, ctx)
		return nil

	}
}

func internalServerError(err interface{}, ctx *fiber.Ctx) bool {
	response, ok := err.(*InternalServerErrorStruct)
	if ok {
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"success": false,
			"message": response.Error(),
		})
		return true
	}
	return false
}

func badRequestError(err interface{}, ctx *fiber.Ctx) bool {
	response, ok := err.(*BadRequestErrorStruct)
	if ok {
		ctx.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"success": false,
			"message": response.Error(),
		})
		return true
	}
	return false
}

func notFoundError(err interface{}, ctx *fiber.Ctx) bool {
	response, ok := err.(*NotFoundErrorStruct)
	if ok {
		ctx.Status(http.StatusNotFound).JSON(map[string]interface{}{
			"success": false,
			"message": response.Error(),
			"data":    map[string]interface{}{},
		})

		return true
	}
	return false
}

func forbiddenError(err interface{}, ctx *fiber.Ctx) bool {
	response, ok := err.(*ForbiddenErrorStruct)
	if ok {
		ctx.Status(http.StatusForbidden).JSON(map[string]interface{}{
			"success": false,
			"message": response.Error(),
		})

		return true
	}
	return false
}

func jwtValidateError(err interface{}, ctx *fiber.Ctx) bool {
	response, ok := err.(*jwt.ValidationError)
	if ok {
		ctx.Status(http.StatusForbidden).JSON(map[string]interface{}{
			"success": false,
			"message": response.Error(),
		})

		return true
	}
	return false
}

func validationError(err interface{}, ctx *fiber.Ctx) bool {
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		var report string
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				field := err.Field()
				report = fmt.Sprintf("%s cannot be null",
					strings.ToLower(field))
			case "email":
				report = fmt.Sprintf("%s is not valid email",
					strings.ToLower(err.Field()))
			case "gte":
				report = fmt.Sprintf("%s value must be greater than %s",
					strings.ToLower(err.Field()), err.Param())
			case "lte":
				report = fmt.Sprintf("%s value must be lower than %s",
					strings.ToLower(err.Field()), err.Param())
			}
		}

		ctx.Status(http.StatusBadRequest).JSON(helpers.FailedResponse(report))

		return true
	}
	return false
}
