package database

import (
	"log"
	usermodel "managerservice/domains/user/model"

	"gorm.io/gorm"
)

func seeder(db *gorm.DB) {
	users := []usermodel.User{
		{
			Name:        "admin",
			Email:       "admin@gmail.com",
			Password:    "$2a$12$ZIgVJrMrQcaJspeW/992z.M5D9qGcauae4hJX2qduYd7xkzfhR0da",
			PhoneNumber: "099879162",
			Gender:      "male",
			Role:        "admin",
		},
		{
			Name:        "user",
			Email:       "user@gmail.com",
			Password:    "$2a$12$3GRupQiAEHqCRtJDwTfetO8eNWjgMK7ej1yBCknj2GpDOtkavvQ82",
			PhoneNumber: "099879162",
			Gender:      "female",
			Role:        "user",
		},
	}
	err := db.Model(&usermodel.User{}).Create(&users).Error
	if err != nil {
		log.Fatal(err)
	}
}
