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
	RegisterUser(dto.Register, *fasthttp.RequestCtx) (models.User, error)
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

func (a *authS) RegisterUser(registerDTO dto.Register, ctx *fasthttp.RequestCtx) (models.User, error) {
	if registerDTO.Username == "" {
		return models.User{}, fmt.Errorf("Register gagal!")
	}
	
	if registerDTO.Password == "" {
		return models.User{}, fmt.Errorf("Register gagal!")
	}

	user, err := a.authR.InsertUser(registerDTO, ctx)
	if err != nil {
		return user, fmt.Errorf("Register gagal!")
	}

	return user, nil
}

