package usercore

import "time"

type Core struct {
	Id          uint
	Name        string
	Email       string
	Password    string
	PhoneNumber string
	Role        string
	Token       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
