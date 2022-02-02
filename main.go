package main

import (
	"github.com/goCrudApp/database"
	"github.com/goCrudApp/routes"
	"github.com/gofiber/fiber/middleware/cors"
	"github.com/gofiber/fiber"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8000")

}
