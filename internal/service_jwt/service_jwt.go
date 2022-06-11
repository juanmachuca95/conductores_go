package service

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/juanmachuca95/conductores_go/domains/models"
)

type JWTService interface {
	GenerateToken(models.User, []*models.Role) string
	ValidateToken(string) (*jwt.Token, error)
	ExtractDataInfoFromJWT(string) (interface{}, error)
}

type claims struct {
	User  models.User `json:"user"`
	Roles []string    `json:"roles"`
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

func (service *jwtServices) GenerateToken(user models.User, roles []*models.Role) string {
	var rols []string
	for _, rol := range roles {
		rols = append(rols, rol.Role)
	}
	claims := &claims{
		Roles: rols,
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

func (service *jwtServices) ExtractDataInfoFromJWT(tokenString string) (interface{}, error) {
	t, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(service.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if t.Valid {
		claims := t.Claims.(jwt.MapClaims)
		return claims["roles"], nil
	}

	return nil, errors.New("No fue posible procesar su token.")
}
