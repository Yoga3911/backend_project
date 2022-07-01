package repository

import (
	"crud/dto"
	"crud/sql"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
)

type AuthR interface {
	CheckUsername(*fasthttp.RequestCtx, dto.Login) pgx.Row
	InsertUser(*fasthttp.RequestCtx, dto.Register) error
}

type authR struct {
	db *pgxpool.Pool
}

func NewAuthR(db *pgxpool.Pool) AuthR {
	return &authR{
		db: db,
	}
}

func (a *authR) CheckUsername(ctx *fasthttp.RequestCtx, loginDTO dto.Login) pgx.Row {
	return a.db.QueryRow(ctx, sql.Authentication, loginDTO.Username)
}

func (a *authR) InsertUser(ctx *fasthttp.RequestCtx, registerDTO dto.Register) error {
	_, err := a.db.Exec(ctx, sql.InsertUser, registerDTO.Id, registerDTO.Username, registerDTO.Email, registerDTO.Password, registerDTO.Address, registerDTO.CreatedAt, registerDTO.UpdatedAt)

	return err
}
