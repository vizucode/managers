package userhandler

import (
	usercore "managerservice/domains/user/core"
	"managerservice/exceptions"
	"managerservice/middlewares"
	"managerservice/utils/helpers"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type activityHandler struct {
	service usercore.IServiceUser
}

var validate = validator.New()

func New(service usercore.IServiceUser) *activityHandler {
	return &activityHandler{
		service: service,
	}
}

func (h *activityHandler) Create(c *fiber.Ctx) error {
	request := Request{}

	err := c.BodyParser(&request)
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	err = validate.Struct(&request)
	if err != nil {
		return err
	}

	h.service.Create(usercore.Core{
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		Name:        request.Name,
		Password:    request.Password,
		Gender:      request.Gender,
	})

	return c.Status(http.StatusCreated).JSON(helpers.SuccessActionResponse("success insert data"))
}

func (h *activityHandler) Update(c *fiber.Ctx) error {
	request := RequestUpdate{}

	err := c.BodyParser(&request)
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	paramID := c.Params("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		panic(exceptions.NewBadRequestError("param not an number"))
	}

	err = validate.Struct(&request)
	if err != nil {
		return err
	}

	h.service.Update(usercore.Core{
		Id:          uint(id),
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		Name:        request.Name,
		Gender:      request.Gender,
	})

	return c.Status(http.StatusOK).JSON(helpers.SuccessActionResponse("success update data"))
}

func (h *activityHandler) Delete(c *fiber.Ctx) error {
	paramID := c.Params("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		panic(exceptions.NewBadRequestError("param not an number"))
	}

	h.service.Delete(usercore.Core{
		Id: uint(id),
	})

	return c.Status(http.StatusOK).JSON(helpers.SuccessActionResponse("success delete data"))
}

func (h *activityHandler) Get(c *fiber.Ctx) error {
	id, _ := middlewares.ExtractToken(c)

	result := h.service.Get(usercore.Core{
		Id: uint(id),
	})

	return c.Status(http.StatusOK).JSON(helpers.SuccessGetResponseData(Response{
		Id:          result.Id,
		Name:        result.Name,
		Email:       result.Email,
		Gender:      result.Gender,
		PhoneNumber: result.PhoneNumber,
	}))
}

func (h *activityHandler) GetAll(c *fiber.Ctx) error {
	results := h.service.GetAll()

	response := []Response{}

	for _, result := range results {
		response = append(response, Response{
			Id:          result.Id,
			Name:        result.Name,
			Email:       result.Email,
			Gender:      result.Gender,
			PhoneNumber: result.PhoneNumber,
		})
	}

	return c.Status(http.StatusOK).JSON(helpers.SuccessGetResponseData(response))
}
