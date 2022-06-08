package repository

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/juanmachuca95/conductores_go/domains/models"
)

type JwtRepository interface {
	GenerateToken(*models.User, []*models.Role) (string, error)
	ValidateToken(string) (*jwt.Token, error)
	ExtractDataInfoFromJWT(string) (interface{}, error)
}
