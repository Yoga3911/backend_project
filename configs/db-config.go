package configs

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

func DatabaseConnection() *pgxpool.Pool {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env file not found!")
	}

	dsn := "dev"
	switch dsn {
	case "dev":
		dsn = fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
			os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	case "prod":
		dsn = os.Getenv("DATABASE_URL")
	}

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal(err)
	}

	config.MaxConns = 20
	config.MinConns = 5
	config.MaxConnIdleTime = 5 * time.Minute
	config.MaxConnLifetime = 60 * time.Minute

	ctx := context.Background()
	pg, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	migration := "true"
	if migration == "true" {
		Migration(pg, ctx)
	} else if migration == "false" {
		Rollback(pg, ctx)
	}

	return pg
}
