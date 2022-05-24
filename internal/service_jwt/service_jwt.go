package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	auth "github.com/juanmachuca95/spaceguru/services/auth/models"
)

type JWTService interface {
	GenerateToken(user auth.User) string
	ValidateToken(token string) (*jwt.Token, error)
}

type claims struct {
	User auth.User `json:"user"`
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

func (service *jwtServices) GenerateToken(user auth.User) string {
	claims := &claims{
		User: user,
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

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token %s", token.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})
}
