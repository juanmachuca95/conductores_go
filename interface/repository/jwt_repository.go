package repository

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/juanmachuca95/conductores_go/domains/models"
)

type JwtRepository interface {
	GenerateToken(models.User, []string) string
	ValidateToken(string) (*jwt.Token, error)
	ExtractDataInfoFromJWT(string) (interface{}, error)
}

type claims struct {
	User  models.User `json:"user"`
	Roles []string    `json:"roles"`
	jwt.StandardClaims
}

type jwtRepository struct {
	secretKey string
}

func NewJwtRepository() JwtRepository {
	return &jwtRepository{secretKey: getSecretKey()}
}

func getSecretKey() string {
	secret := os.Getenv("TOKEN_API_AUTH")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (j *jwtRepository) GenerateToken(user models.User, roles []string) string {
	claims := &claims{
		Roles: roles,
		User:  user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	_token, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return _token
}

func (j *jwtRepository) ValidateToken(receivedToken string) (*jwt.Token, error) {
	return jwt.Parse(receivedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token")
		}
		return []byte(j.secretKey), nil
	})
}

func (j *jwtRepository) ExtractDataInfoFromJWT(tokenString string) (interface{}, error) {
	t, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})

	if err != nil {
		panic(err.Error())
	}

	if t.Valid {
		claims := t.Claims.(jwt.MapClaims)
		return claims["roles"], nil
	}

	return nil, errors.New("No fue posible procesar su token.")
}