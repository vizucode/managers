package userhandler

type Response struct {
	Id          uint   `json:"id"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
}
