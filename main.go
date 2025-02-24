package main

import (
	routes "yakz/Routes"
	"yakz/config"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	config.ConfigDataBase()
	_ = godotenv.Load()
	app := fiber.New()

	routes.SetupUserApi(app)
	routes.SetupBookApi(app)
	routes.SetupCartApi(app)

	app.Listen(":8080")
}
