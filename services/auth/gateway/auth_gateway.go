package auth

import (
	"database/sql"

	m "github.com/juanmachuca95/conductores_go/services/auth/models"
)

type AuthGateway interface {
	Login(*m.Login) (*m.UserToken, error)
	Register(*m.Register) (*m.UserToken, error)
	GetUser(id int64) (*m.User, error)
	CreateRoleUser(role string, id int64) (*int64, error)
	GetRolesUser(id int64) (*[]string, error)
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

func (s *AuthInDB) GetUser(id int64) (*m.User, error) {
	return s.getUser(id)
}

func (s *AuthInDB) CreateRoleUser(role string, id int64) (*int64, error) {
	return s.createRoleUser(role, id)
}

func (s *AuthInDB) GetRolesUser(id int64) (*[]string, error) {
	return s.getRolesUser(id)
}

func NewAuthGateway(db *sql.DB) AuthGateway {
	return &AuthInDB{NewAuthStorageGateway(db)}
}
