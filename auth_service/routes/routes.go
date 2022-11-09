package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	authhandler "authentication/domains/auth/handler"
	authrepo "authentication/domains/auth/repository"
	authservice "authentication/domains/auth/service"
)

func InitRoutes(ctx *fiber.App, db *gorm.DB) {
	/*
		Dependency Injection
	*/

	authRepo := authrepo.New(db)
	authService := authservice.New(authRepo)
	authHandler := authhandler.New(authService)

	/*
		Routes
	*/

	ctx.Post("/login", authHandler.Login)
}
