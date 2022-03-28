package main

import (
	"log"

	"github.com/dakicka/tradingApp/api/api"
	"github.com/dakicka/tradingApp/api/config"
	"github.com/dakicka/tradingApp/api/database"
	"github.com/dakicka/tradingApp/api/repository"
	"github.com/dakicka/tradingApp/api/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

//TODO:
// - Further abstract token validation and user parsing from token
// - Fieldvalidations (https://medium.com/@apzuk3/input-validation-in-golang-bc24cdec1835)
// - Consistency in error handling!

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

	d := database.New(config)
	d.Migrate()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// Setup all Routes
	service := service.New(&repository.UsersSQL{})
	api.NewUser(app, service)
	//routes.Setup(app)

	log.Fatal(app.Listen(":8080"))
}
