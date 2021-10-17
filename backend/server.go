package main

import (
	"appoiment-backend/database"
	"appoiment-backend/middlewares"
	"appoiment-backend/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
		panic("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")


	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable.")
	}

	database.Connect(uri)
	defer database.Disconnect()

	app := fiber.New()
	routes.UserSetup(app)

	app.Use(middlewares.AuthMiddleware)

	routes.ClientSetup(app)
	routes.ServiceSetup(app)
	
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

    app.Listen(":3000")
}