package usercore

import "time"

type Core struct {
	Id          uint
	Email       string
	Password    string
	Name        string
	PhoneNumber string
	Gender      string
	Role        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
