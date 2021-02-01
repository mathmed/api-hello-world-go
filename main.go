package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/mathmed/api-hello-world-go/src/routes"
)


func main() {
	godotenv.Load()
	app := fiber.New()
	app.Use(logger.New())
	routes.SetupRoutes(app)
	port := os.Getenv("PORT")
	
	if port == "" {
		port = "8080"
	}

	err := app.Listen(":" + port)

	if err != nil {
		panic(err)
	}

	log.Printf("Listening on port %s", port)

}
