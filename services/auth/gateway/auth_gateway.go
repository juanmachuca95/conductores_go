package auth

import (
	database "github.com/juanmachuca95/spaceguru/internal/databases"
	service_jwt "github.com/juanmachuca95/spaceguru/internal/service_jwt"
	m "github.com/juanmachuca95/spaceguru/services/auth/models"
)

type AuthGateway interface {
	Login(*m.Login) (*m.UserToken, error)
	Register(*m.Register) (*m.UserToken, error)
}

func NewAuthService() AuthGateway {
	return &AuthService{database.NewMySQLClient(), service_jwt.JWTAuthService()}
}

func (s *AuthService) Login(u *m.Login) (*m.UserToken, error) {
	return s.login(u)
}

func (s *AuthService) Register(u *m.Register) (*m.UserToken, error) {
	return s.register(u)
}
