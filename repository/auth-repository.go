package repository

import "github.com/jackc/pgx/v4/pgxpool"

type AuthR interface {
}

type authR struct {
	db *pgxpool.Pool
}

func NewAuthR(db *pgxpool.Pool) AuthR {
	return &authR{
		db: db,
	}
}
