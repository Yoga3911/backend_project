package main

import (
	"context"
	"crud/configs"
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
		log.Fatal(".env not found")
	}

	if len(os.Args[1:]) != 0 {
		ctx := context.Background()
		switch os.Args[1] {
		case "-m":
			configs.Migration(routes.DB, ctx)
		case "-r":
			configs.Rollback(routes.DB, ctx)
		}
		return
	}

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
