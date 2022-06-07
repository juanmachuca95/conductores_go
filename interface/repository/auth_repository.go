package repository

import (
	"database/sql"
	"errors"

	"github.com/juanmachuca95/conductores_go/domains/models"
	"github.com/juanmachuca95/conductores_go/usescases/repository"
	"golang.org/x/crypto/bcrypt"
)

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) repository.AuthRepository {
	return &authRepository{db: db}
}

func (a *authRepository) Login(l *models.Login) (*models.User, error) {
	stmt, err := a.db.Prepare("SELECT id, name, email, created_at, updated_at FROM users WHERE email = ?;")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	var user models.User
	var password string
	err = stmt.QueryRow(l.Email).Scan(&user.Id, &user.Name, &user.Email, &password, &user.Created_at, &user.Updated_at)
	if err != nil {
		panic(err)
	}

	hashByte := []byte(password)
	passwordByte := []byte(l.Password)
	err = bcrypt.CompareHashAndPassword(hashByte, passwordByte) // Validación de la contrasenia por el hash
	if err != nil {
		return &models.User{}, errors.New("Credenciales incorrectas.")
	}

	return &user, nil
}
