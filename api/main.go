package main

import (
	"log"

	"github.com/dakicka/tradingApp/api/database"
	"github.com/dakicka/tradingApp/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

//TODO:
// Abstract token validation and user parsing from token

func main() {
	// Get .env variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// Setup all Routes
	routes.Setup(app)

	log.Fatal(app.Listen(":3000"))
}
