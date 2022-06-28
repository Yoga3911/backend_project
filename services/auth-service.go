package services

import (
	"crud/dto"
	"crud/models"
	"crud/repository"
	"fmt"

	"github.com/valyala/fasthttp"
)

type AuthS interface {
	LoginUser(dto.Login, *fasthttp.RequestCtx) (models.User, error)
	RegisterUser()
}

type authS struct {
	authR repository.AuthR
}

func NewAuthS(authR repository.AuthR) AuthS {
	return &authS{
		authR: authR,
	}
}

func (a *authS) LoginUser(loginDTO dto.Login, ctx *fasthttp.RequestCtx) (models.User, error) {
	user, err := a.authR.CheckUsernamePassword(loginDTO, ctx)
	if err != nil {
		return user, fmt.Errorf("Username atau Password salah!")
	}

	return user, nil
}

func (a *authS) RegisterUser() {

}
