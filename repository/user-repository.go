package repository

import (
	"crud/models"
	"crud/sql"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/valyala/fasthttp"
)

type UserR interface {
	GetById(*fasthttp.RequestCtx, string) (models.User, error)
}

type userR struct {
	db *pgxpool.Pool
}

func NewUserR(db *pgxpool.Pool) UserR {
	return &userR{
		db: db,
	}
}

func (u *userR) GetById(ctx *fasthttp.RequestCtx, userId string) (models.User, error) {
	var user models.User
	
	err := u.db.QueryRow(ctx, sql.GetUserById, userId).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Address, &user.RoleId, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}
