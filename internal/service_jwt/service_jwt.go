package service

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	auth "github.com/juanmachuca95/spaceguru/services/auth/models"
)

type JWTService interface {
	GenerateToken(auth.User, []auth.RolesUser) string
	ValidateToken(string) (*jwt.Token, error)
	ExtractDataInfoFromJWT(string) ([]auth.RolesUser, error)
}

type claims struct {
	User  auth.User        `json:"user"`
	Roles []auth.RolesUser `json:"roles"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
}

func JWTAuthService() JWTService {
	return &jwtServices{secretKey: getSecretKey()}
}

func getSecretKey() string {
	secret := os.Getenv("TOKEN_API_AUTH")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtServices) GenerateToken(user auth.User, roles []auth.RolesUser) string {
	claims := &claims{
		Roles: roles,
		User:  user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	_token, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return _token
}

func (service *jwtServices) ValidateToken(receivedToken string) (*jwt.Token, error) {
	return jwt.Parse(receivedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token")
		}
		return []byte(service.secretKey), nil
	})
}

func (service *jwtServices) ExtractDataInfoFromJWT(tokenString string) ([]auth.RolesUser, error) {
	t, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(service.secretKey), nil
	})

	if err != nil {
		panic(err.Error())
	}

	if t.Valid {
		claims := t.Claims.(jwt.MapClaims)

		return claims["roles"].([]auth.RolesUser), nil
	}

	return []auth.RolesUser{}, errors.New("No fue posible procesar su token.")
}
