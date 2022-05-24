package auth

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	database "github.com/juanmachuca95/spaceguru/internal/databases"
	m "github.com/juanmachuca95/spaceguru/services/auth/models"
	q "github.com/juanmachuca95/spaceguru/services/auth/querys"
	"golang.org/x/crypto/bcrypt"
)

type AuthStorage interface {
	login(u *m.Login) (string, error)
	register(u *m.Register) (*m.UserToken, error)
}

type AuthService struct {
	*database.MySQLClient
}

func NewAuthStorageGateway(db *database.MySQLClient) AuthGateway {
	return &AuthService{db}
}

type claims struct {
	User m.User `json:"user"`
	jwt.StandardClaims
}

func NewClaims(user m.User) *claims {
	return &claims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			Subject:   "1",
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
	}
}

func (s *AuthService) login(u *m.Login) (*m.UserToken, error) {
	stmt, err := s.Prepare(q.GetUser())
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	var user m.User
	var password string
	err = stmt.QueryRow(u.Email).Scan(&user.Id, &user.Name, &user.Email, &password, &user.Rol)
	if err != nil {
		return &m.UserToken{}, errors.New("Este usuario no existe.")
	}

	hashByte := []byte(password)
	passwordByte := []byte(u.Password)
	err = bcrypt.CompareHashAndPassword(hashByte, passwordByte) // Validación de la contrasenia por el hash
	if err != nil {
		return &m.UserToken{}, errors.New("Credenciales incorrectas.")
	}

	claims := NewClaims(user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := os.Getenv("TOKEN_API_AUTH")
	_token, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return &m.UserToken{}, err
	}

	return &m.UserToken{
		Token: _token,
	}, nil
}

func (s *AuthService) register(u *m.Register) (*m.UserToken, error) {
	stmt, err := s.Prepare(q.InsertUser())
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	u.Rol = os.Getenv("DEFAULT_ROL")
	hashPass, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	var passwordHashed string = string(hashPass)
	result, err := stmt.Exec(u.Name, u.Email, passwordHashed, u.Rol)
	if err != nil {
		return &m.UserToken{}, err
	}

	users_id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	user, _ := s.getUser(users_id)
	claims := NewClaims(*user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := os.Getenv("TOKEN_API_AUTH")
	_token, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return &m.UserToken{}, err
	}

	return &m.UserToken{
		Token: _token,
	}, nil
}

func (s *AuthService) getUser(id int64) (*m.User, error) {
	stmt, err := s.Prepare(q.GetUserById())
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	var user m.User
	err = stmt.QueryRow(id).Scan(&user.Id, &user.Name, &user.Email, &user.Rol)
	if err != nil {
		panic(err.Error())
	}

	return &user, nil
}
