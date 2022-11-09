package usermodel

import (
	usercore "authentication/domains/auth/core"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string
	Email       string
	Password    string
	PhoneNumber string
	Role        string
}

func ToCore(model User) usercore.Core {
	return usercore.Core{
		Email:       model.Email,
		Name:        model.Name,
		PhoneNumber: model.PhoneNumber,
		Role:        model.Role,
		Password:    model.Password,
	}
}

func ToModel(Core usercore.Core) User {
	return User{
		Email:       Core.Email,
		Name:        Core.Name,
		PhoneNumber: Core.PhoneNumber,
		Role:        Core.Role,
		Password:    Core.Password,
	}
}
