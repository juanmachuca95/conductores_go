package auth

import (
	"errors"

	database "github.com/juanmachuca95/spaceguru/internal/databases"
	m "github.com/juanmachuca95/spaceguru/services/auth/models"
	q "github.com/juanmachuca95/spaceguru/services/auth/querys"
)

type AuthStorage interface {
	login(u *m.Login) (string, error)
}

type AuthService struct {
	*database.MySQLClient
}

func NewAuthStorageGateway(db *database.MySQLClient) AuthGateway {
	return &AuthService{db}
}

func (s *AuthService) login(u *m.Login) (string, error) {
	stmt, err := s.Prepare(q.GetUser())
	if err != nil {
		return "", errors.New("Este usuario no existe.")
	}
	defer stmt.Close()

	return "", nil
}
