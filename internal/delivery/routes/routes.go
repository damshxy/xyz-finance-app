package routes

import "github.com/gofiber/fiber/v2"

func RoutesInit(app *fiber.App) {
	api := app.Group("/api")

	consumerRoutes(api)
	transactionRoutes(api)
}