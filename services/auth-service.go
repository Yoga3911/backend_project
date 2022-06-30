package services

import (
	"crud/dto"
	"crud/models"
	"crud/repository"
	"crud/utils"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

type AuthS interface {
	LoginUser(dto.Login, *fasthttp.RequestCtx) (models.User, error)
	RegisterUser(dto.Register, *fasthttp.RequestCtx) error
}

type authS struct {
	authR repository.AuthR
	jwtS  JWTService
}

func NewAuthS(authR repository.AuthR, jwtS JWTService) AuthS {
	return &authS{
		authR: authR,
		jwtS:  jwtS,
	}
}

func (a *authS) LoginUser(loginDTO dto.Login, ctx *fasthttp.RequestCtx) (models.User, error) {
	user, err := a.authR.CheckUsername(loginDTO, ctx)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return user, fmt.Errorf("Username atau Password salah!")
		}
		
		return user, err
	}

	err = utils.ComparePassword(loginDTO.Password, user.Password)
	if err != nil {
		return user, fmt.Errorf("Username atau Password salah!")
	}

	user.Token = a.jwtS.GenerateToken(user)

	return user, nil
}

func (a *authS) RegisterUser(registerDTO dto.Register, ctx *fasthttp.RequestCtx) error {
	hash, err := utils.HashAndSalt(registerDTO.Password)
	if err != nil {
		return err
	}
	registerDTO.Password = hash

	registerDTO.Id = uuid.New().String()

	timeMili := time.Now().UnixMilli()
	registerDTO.CreatedAt = timeMili
	registerDTO.UpdatedAt = timeMili

	err = a.authR.InsertUser(registerDTO, ctx)
	if err != nil {
		if strings.Contains(err.Error(), "username") {
			return fmt.Errorf("Username telah terdaftar!")
		} else if strings.Contains(err.Error(), "email") {
			return fmt.Errorf("Email telah terdaftar!")
		} else {
			return fmt.Errorf(err.Error())
		}
	}

	return nil
}
