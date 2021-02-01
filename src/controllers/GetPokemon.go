package controllers

import (
	"api-hello-world-go/src/models"
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetPokemon(context *fiber.Ctx) error {
	response, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + context.Params("pokemon"))

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
			"code": 500,
		})
	}

	if response.StatusCode != http.StatusOK {
		return context.Status(response.StatusCode).JSON(fiber.Map{
			"message": response,
			"code": 400,
		})
	}

	var pokemon models.Pokemon


	err = json.NewDecoder(response.Body).Decode(&pokemon)

	if err != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
			"code": 500,
		})
	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Ok",
		"code": 200,
		"data": models.Pokemon{
			Id: pokemon.Id,
			Weight: pokemon.Weight,
			Name: pokemon.Name,
			Base_experience: pokemon.Base_experience,
			Sprites: pokemon.Sprites,
		},
	})
}