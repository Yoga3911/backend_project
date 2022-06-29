package repository

import (
	"crud/dto"
	"crud/models"
	"crud/sql"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
)

type AuthR interface {
	CheckUsernamePassword(dto.Login, *fasthttp.RequestCtx) (models.User, error)
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

func (a *authR) CheckUsernamePassword(loginDTO dto.Login, ctx *fasthttp.RequestCtx) (models.User, error) {
	var user models.User
	err := a.db.QueryRow(ctx, sql.Authentication, loginDTO.Username, loginDTO.Password).
		Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Address)

	return user, err
}

func (a *authR) InsertUser(registerDTO dto.Register, ctx *fasthttp.RequestCtx) error {
	_, err := a.db.Exec(ctx, sql.InsertUser, registerDTO.Username, registerDTO.Email, registerDTO.Password, registerDTO.Address)

	if strings.Contains(err.Error(), "username") {
		err = fmt.Errorf("Username telah terdaftar!")
	} else if strings.Contains(err.Error(), "email") {
		err = fmt.Errorf("Email telah terdaftar!")
	}

	return err
}
