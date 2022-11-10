package usermodel

import (
	usercore "managerservice/domains/user/core"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string
	Password    string
	Name        string
	PhoneNumber string
	Gender      string
	Role        string
}

func ToCore(model User) usercore.Core {
	return usercore.Core{
		Id:          model.ID,
		Email:       model.Email,
		Name:        model.Name,
		PhoneNumber: model.PhoneNumber,
		Password:    model.Password,
		Gender:      model.Gender,
		Role:        "member",
	}
}

func ToModel(Core usercore.Core) User {
	return User{
		Email:       Core.Email,
		Name:        Core.Name,
		PhoneNumber: Core.PhoneNumber,
		Password:    Core.Password,
		Gender:      Core.Gender,
		Role:        Core.Role,
	}
}
