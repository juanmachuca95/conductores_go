package auth

import (
	database "github.com/juanmachuca95/spaceguru/internal/databases"
	m "github.com/juanmachuca95/spaceguru/services/auth/models"
)

type AuthGateway interface {
	Login(*m.Login) (*m.User, error)
	Register(*m.Register) (*m.User, error)
}

type AuthInDB struct {
	AuthStorage
}

func NewAuthService() AuthGateway {
	return &AuthService{database.NewMySQLClient()}
}

func (s *AuthService) Login(u *m.Login) (*m.User, error) {
	return s.login(u)
}

func (s *AuthService) Register(u *m.Register) (*m.User, error) {
	return s.Register(u)
}
