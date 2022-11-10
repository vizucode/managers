package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	userhandler "managerservice/domains/user/handler"
	userrepo "managerservice/domains/user/repository"
	userservice "managerservice/domains/user/service"
	"managerservice/middlewares"
)

func InitRoutes(ctx *fiber.App, db *gorm.DB) {
	/*
		Dependency Injection
	*/

	userRepo := userrepo.New(db)
	userService := userservice.New(userRepo)
	userHandler := userhandler.New(userService)

	/*
		Routes
	*/

	ctx.Post("/register", userHandler.Create)
	ctx.Put("/verify", userHandler.Verify)
}
