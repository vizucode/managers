package userhandler

type Response struct {
	Id          uint   `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
	Role        string `json:"role"`
}
