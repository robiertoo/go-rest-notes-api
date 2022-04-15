package main

import (
	"rest-api/database"

	"github.com/gofiber/fiber/v2"

	"rest-api/router"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	router.SetupRoutes(app)

	app.Listen(":3000")
}
