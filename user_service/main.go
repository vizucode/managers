package main

import (
	"managerservice/config"
	"managerservice/exceptions"
	"managerservice/routes"
	utils "managerservice/utils/database/postgre"

	jsongoccy "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	cfg := config.GetConfig()
	db := utils.InitDB(cfg)
	app := fiber.New(fiber.Config{
		JSONEncoder:  jsongoccy.Marshal,
		JSONDecoder:  jsongoccy.Unmarshal,
		ErrorHandler: exceptions.CustomErrorHandling,
	})
	app.Use(recover.New())

	routes.InitRoutes(app, db)

	if err := app.Listen(":" + cfg.SERVER_PORT); err != nil {
		panic(err)
	}
}
