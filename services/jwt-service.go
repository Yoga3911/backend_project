package services

import (
	"crud/dto"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type JWTService interface {
	GenerateToken(user dto.UserLogin) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtSevice struct {
	issuer    string
	secretKey string
}

type jwtCustomClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	RoleId   int8   `json:"role_id"`
	jwt.StandardClaims
}

func NewJWTService() JWTService {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	return &jwtSevice{
		issuer:    "myproject",
		secretKey: os.Getenv("SECRET_KEY"),
	}
}

func (j *jwtSevice) GenerateToken(user dto.UserLogin) string {
	claims := jwtCustomClaim{
		Username: user.Username,
		Email:    user.Email,
		Address:  user.Address,
		RoleId:   user.RoleId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 0, 1).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		log.Println(err)
	}

	return t
}

func (j *jwtSevice) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}
