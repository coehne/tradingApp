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
// - Further abstract token validation and user parsing from token
// - Fieldvalidations (https://medium.com/@apzuk3/input-validation-in-golang-bc24cdec1835)

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
