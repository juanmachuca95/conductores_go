package auth

import (
	"errors"

	database "github.com/juanmachuca95/spaceguru/internal/databases"
	m "github.com/juanmachuca95/spaceguru/services/auth/models"
	q "github.com/juanmachuca95/spaceguru/services/auth/querys"
	"golang.org/x/crypto/bcrypt"
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

func (s *AuthService) login(u *m.Login) (*m.User, error) {
	stmt, err := s.Prepare(q.GetUser())
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	var user m.User
	err = stmt.QueryRow(u.Email).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return &m.User{}, errors.New("No existe el usuario")
	}

	hashByte := []byte(user.Password)
	passwordByte := []byte(u.Password)
	err = bcrypt.CompareHashAndPassword(hashByte, passwordByte) // Validación de la contrasenia por el hash
	if err != nil {
		return &m.User{}, errors.New("Credenciales incorrectas.")
	}

	return &user, nil
}
