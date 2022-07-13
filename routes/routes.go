package routes

import (
	"crud/configs"
	"crud/controllers"
	"crud/repository"
	"crud/services"
	"crud/utils"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

var (
	DB  *pgxpool.Pool       = configs.DatabaseConnection()
	JWT services.JWTService = services.NewJWTService()

	Firebase services.Storage = services.NewStorage()

	authR repository.AuthR  = repository.NewAuthR(DB)
	authS services.AuthS    = services.NewAuthS(authR, JWT)
	authC controllers.AuthC = controllers.NewAuthC(authS)

	userR repository.UserR  = repository.NewUserR(DB)
	userS services.UserS    = services.NewUserS(userR)
	userC controllers.UserC = controllers.NewUserC(userS)

	productR repository.ProductR  = repository.NewProductR(DB)
	productS services.ProductS    = services.NewProductS(productR)
	productC controllers.ProductC = controllers.NewProductC(productS)

	fileC controllers.FileC = controllers.NewFileC(Firebase)
)

func Data(app *fiber.App) {
	app.Use(cors.New())
	app.Get("/", OK)

	api := app.Group("/api/v1")

	api.Post("/auth/login", authC.Login)
	api.Post("/auth/register", authC.Register)
	api.Get("/product", productC.GetAllProduct)
	api.Get("/product/:productId", productC.GetProductById)
	api.Post("/image/upload", fileC.Upload)

	// API KEY
	api.Use(func(c *fiber.Ctx) error {
		err := godotenv.Load()
		if err != nil {
			log.Println(err)
		}

		key := c.Get("API-KEY")
		if key == "" {
			return utils.Response(c, 401, nil, "API KEY tidak ditemukan", false)
		}

		if key != os.Getenv("API_KEY") {
			return utils.Response(c, 401, nil, "API KEY tidak sesuai", false)
		}

		return c.Next()
	})

	// Middleware - Check Token
	api.Use(func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return utils.Response(c, 401, nil, "Token tidak ditemukan", false)
		}

		_, err := JWT.ValidateToken(token)
		if err != nil {
			return utils.Response(c, 401, nil, "Token tidak valid", false)
		}

		return c.Next()
	})

	api.Get("/user/:id", userC.GetUserById)

	// Middleware - Check Token
	api.Use(func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		t, _ := JWT.ValidateToken(token)

		claims := t.Claims.(jwt.MapClaims)

		if (claims["role_id"]) == float64(2) {
			return c.Next()
		}

		return utils.Response(c, 403, nil, "You are not a seller!", false)
	})

	api.Post("/product", productC.InsertProduct)
	api.Put("/product", productC.EditProduct)
	api.Delete("/product", productC.DeleteProduct)
}

func OK(c *fiber.Ctx) error {
	return utils.Response(c, 200, nil, "OK!", true)
}
