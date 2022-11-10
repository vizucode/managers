package userhandler

type Request struct {
	Email       string `json:"email" form:"email" validate:"required"`
	Name        string `json:"name" form:"name" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required"`
	Gender      string `json:"gender" form:"gender" validate:"required"`
	Password    string `json:"password" form:"password" validate:"required"`
}

type RequestUpdate struct {
	Email       string `json:"email" form:"email"`
	Name        string `json:"name" form:"name"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Gender      string `json:"gender" form:"gender"`
}
