package configs

import (
	"context"
	"crud/sql"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func Migration(db *pgxpool.Pool, ctx context.Context) {
	_, err := db.Exec(ctx, sql.Migration)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(ctx, sql.CallMigration)
	if err != nil {
		panic(err)
	}

	log.Println("Migration success!")
}
