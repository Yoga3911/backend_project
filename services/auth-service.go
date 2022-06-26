package services

import "crud/repository"

type AuthS interface {
	LoginUser()
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

func (a *authS) LoginUser() {

}

func (a *authS) RegisterUser() {

}
