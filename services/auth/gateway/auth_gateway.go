package auth

import (
	database "github.com/juanmachuca95/spaceguru/internal/databases"
	m "github.com/juanmachuca95/spaceguru/services/auth/models"
)

type AuthGateway interface {
	Login(*m.Login) (string, error)
	Register(*m.Register) (string, error)
}

type AuthInDB struct {
	AuthStorage
}

func NewAuthService() AuthGateway {
	return &AuthService{database.NewMySQLClient()}
}

func (s *AuthService) Login(u *m.Login) (string, error) {
	return s.login(u)
}

func (s *AuthService) Register(u *m.Register) (string, error) {
	return s.Register(u)
}
