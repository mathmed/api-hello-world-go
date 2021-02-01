package routes

import (
	"api-hello-world-go/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func PokemonRoute(router fiber.Router) {
	router.Get(":pokemon", controllers.GetPokemon)
}