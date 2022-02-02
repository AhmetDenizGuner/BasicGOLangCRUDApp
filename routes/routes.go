package routes

import (
	"github.com/goCrudApp/controllers"
	"github.com/gofiber/fiber"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("api/login", controllers.Login)
}
