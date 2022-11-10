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
	ctx.Use(middlewares.JWTMiddleware)

	ctx.Get("/user", userHandler.Get)

	ctx.Use(middlewares.IsAdmin)

	ctx.Post("/users", userHandler.Create)
	ctx.Put("/user/:id", userHandler.Update)
	ctx.Delete("/user/:id", userHandler.Delete)
	ctx.Get("/users", userHandler.GetAll)
}
