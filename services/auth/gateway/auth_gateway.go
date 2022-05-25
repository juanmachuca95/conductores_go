package auth

import (
	"database/sql"

	m "github.com/juanmachuca95/spaceguru/services/auth/models"
)

type AuthGateway interface {
	Login(*m.Login) (*m.UserToken, error)
	Register(*m.Register) (*m.UserToken, error)
}

type AuthInDB struct {
	AuthStorage
}

func (s *AuthInDB) Login(u *m.Login) (*m.UserToken, error) {
	return s.login(u)
}

func (s *AuthInDB) Register(u *m.Register) (*m.UserToken, error) {
	return s.register(u)
}

func NewAuthGateway(db *sql.DB) AuthGateway {
	return &AuthInDB{NewAuthStorageGateway(db)}
}
