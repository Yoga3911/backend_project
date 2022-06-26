package main

import (
	"crud/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	defer routes.DB.Close()

	app := fiber.New()
	routes.Data(app)

	if err := godotenv.Load(); err != nil {
		log.Fatalln(".env not found")
	}

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
