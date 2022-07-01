package repository

import (
	"crud/sql"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
)

type UserR interface {
	GetById(*fasthttp.RequestCtx, string) pgx.Row
}

type userR struct {
	db *pgxpool.Pool
}

func NewUserR(db *pgxpool.Pool) UserR {
	return &userR{
		db: db,
	}
}

func (u *userR) GetById(ctx *fasthttp.RequestCtx, userId string) pgx.Row {
	return u.db.QueryRow(ctx, sql.GetUserById, userId)
}
