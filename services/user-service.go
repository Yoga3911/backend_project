package services

import (
	"crud/models"
	"crud/repository"
	"fmt"

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
	user, err := u.userR.GetById(ctx, userId)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return user, fmt.Errorf("User tidak ditemukan!")
		}

		return user, err
	}

	return user, nil
}
