package services

import (
	"crud/dto"
	"crud/repository"
	"crud/utils"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

type AuthS interface {
	LoginUser(*fasthttp.RequestCtx, dto.Login) (dto.UserLogin, error)
	RegisterUser(*fasthttp.RequestCtx, dto.Register) error
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

func (a *authS) LoginUser(ctx *fasthttp.RequestCtx, loginDTO dto.Login) (dto.UserLogin, error) {
	var user dto.UserLogin

	err := a.authR.CheckUsername(ctx, loginDTO).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Address, &user.RoleId, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return user, errors.New("username atau password salah")
		}

		return user, err
	}

	err = utils.ComparePassword(loginDTO.Password, user.Password)
	if err != nil {
		return user, errors.New("username atau password salah")
	}

	user.Token = a.jwtS.GenerateToken(user)

	return user, nil
}

func (a *authS) RegisterUser(ctx *fasthttp.RequestCtx, registerDTO dto.Register) error {
	hash, err := utils.HashAndSalt(registerDTO.Password)
	if err != nil {
		return err
	}
	registerDTO.Password = hash

	registerDTO.Id = uuid.New().String()

	timeMili := time.Now().UnixMilli()
	registerDTO.CreatedAt = timeMili
	registerDTO.UpdatedAt = timeMili

	err = a.authR.InsertUser(ctx, registerDTO)
	if err != nil {
		if strings.Contains(err.Error(), "username") {
			return errors.New("username telah terdaftar")
		} else if strings.Contains(err.Error(), "email") {
			return errors.New("email telah terdaftar")
		}

		return fmt.Errorf(err.Error())
	}

	return nil
}
