package main

import (
	"log"

	"github.com/dakicka/tradingApp/api/api"
	"github.com/dakicka/tradingApp/api/config"
	"github.com/dakicka/tradingApp/api/db"
	"github.com/dakicka/tradingApp/api/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

//TODO:
// - Further abstract token validation and user parsing from token
// - Fieldvalidations (https://medium.com/@apzuk3/input-validation-in-golang-bc24cdec1835)
// - Rethink service initialization
// - Find better function names for endpoint, service and repo functions
// - Get rid of .env file and use config instead
// - Add middleware to secure private routes

//Questions:
// - How to improve logging. What fmt or logging or error functions are the right ones?
// - Where to put helperfunctions which are used in the service. i. E. getCashByUserId or check if user
//   has enough stonks or cash to do the trade?
// - What are good guides to improve overall error handling? When to throw which error to what output?
// - Should the service or the endpoint get params from url or user from cookie/token?

func main() {
	// Get .env variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// load config
	config, err := config.New()
	if err != nil {
		log.Fatalf("could not load config: %s", err)
	}

	// Start the databse
	d := db.New(config)
	d.Migrate()

	// Initiate fiber app
	app := fiber.New()
	// Enable cookie support
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// Setup all Routes
	service := service.Init(d)
	api.NewUser(app, service)
	api.NewTrade(app, service)
	api.NewTransaction(app, service)
	api.NewQuote(app, service)

	log.Fatal(app.Listen(":8080"))
}
