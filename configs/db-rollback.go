package configs

import (
	"context"
	"crud/sql"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func Rollback(db *pgxpool.Pool, ctx context.Context) {
	_, err := db.Exec(ctx, sql.Rollback)
	if err != nil {
		log.Println(err)
	}

	_, err = db.Exec(ctx, sql.CallRollback)
	if err != nil {
		log.Println(err)
	}

	log.Println("Rollback success!")
}
