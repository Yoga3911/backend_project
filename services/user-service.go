package services

import (
	"crud/models"
	"crud/repository"
	"errors"

	"github.com/valyala/fasthttp"
)

type UserS interface {
	GetUserById(*fasthttp.RequestCtx, string) (models.User, error)
}

type userS struct {
	userR repository.UserR
}

func NewUserS(userR repository.UserR) UserS {
	return &userS{
		userR: userR,
	}
}

func (u *userS) GetUserById(ctx *fasthttp.RequestCtx, userId string) (models.User, error) {
	var user models.User

	err := u.userR.GetById(ctx, userId).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Address, &user.RoleId, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return user, errors.New("user tidak ditemukan")
		}

		return user, err
	}

	return user, nil
}
