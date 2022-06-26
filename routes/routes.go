package routes

import (
	"crud/configs"
	"crud/controllers"
	"crud/repository"
	"crud/services"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	DB *pgxpool.Pool = configs.DatabaseConnection()

	authR repository.AuthR  = repository.NewAuthR(DB)
	authS services.AuthS    = services.NewAuthS(authR)
	authC controllers.AuthC = controllers.NewAuthC(authS)
)

func Data(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Post("/auth/login", authC.Login)
	api.Post("/auth/register", authC.Register)
}
