package routes

import (
	"github.com/Kosodaka/auth-service/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Post("/api/get-tokens", controllers.GetTokens)
	app.Post("/api/update-tokens", controllers.UpdateTokens)

}
