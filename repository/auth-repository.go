package repository

import (
	"crud/dto"
	"crud/sql"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
)

type AuthR interface {
	CheckUsername(dto.Login, *fasthttp.RequestCtx) pgx.Row
	InsertUser(dto.Register, *fasthttp.RequestCtx) error
}

type authR struct {
	db *pgxpool.Pool
}

func NewAuthR(db *pgxpool.Pool) AuthR {
	return &authR{
		db: db,
	}
}

func (a *authR) CheckUsername(loginDTO dto.Login, ctx *fasthttp.RequestCtx) pgx.Row {
	return a.db.QueryRow(ctx, sql.Authentication, loginDTO.Username)
}

func (a *authR) InsertUser(registerDTO dto.Register, ctx *fasthttp.RequestCtx) error {
	_, err := a.db.Exec(ctx, sql.InsertUser, registerDTO.Id, registerDTO.Username, registerDTO.Email, registerDTO.Password, registerDTO.Address, registerDTO.CreatedAt, registerDTO.UpdatedAt)

	return err
}
