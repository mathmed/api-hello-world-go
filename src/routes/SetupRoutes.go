package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(context *fiber.Ctx) error {
		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Use /api/pokemon/{name} or /api/pokemon/{pokedexNumber} to request pokemon data",
			"code" : 200,
		})
	})

	api := app.Group("/api")

	api.Get("/", func(context *fiber.Ctx) error {
		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Use /api/pokemon/{name} or /api/pokemon/{pokedexNumber} to request pokemon data",
			"code" : 200,
		})
	})

	PokemonRoute(api.Group("/pokemon"))
}