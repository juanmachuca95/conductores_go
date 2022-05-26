package auth

import (
	"database/sql"
	"errors"

	srvJWT "github.com/juanmachuca95/spaceguru/internal/service_jwt"
	m "github.com/juanmachuca95/spaceguru/services/auth/models"
	q "github.com/juanmachuca95/spaceguru/services/auth/querys"
	"golang.org/x/crypto/bcrypt"
)

type AuthStorage interface {
	login(u *m.Login) (*m.UserToken, error)
	register(u *m.Register) (*m.UserToken, error)
	getUser(id int64) (*m.User, error)
	createRoleUser(role string, id int64) (*int64, error)
	getRolesUser(id int64) (*[]string, error)
}

type AuthService struct {
	*sql.DB
	srvJWT srvJWT.JWTService
}

func NewAuthStorageGateway(db *sql.DB) AuthStorage {
	return &AuthService{db, srvJWT.JWTAuthService()}
}

func (s *AuthService) login(u *m.Login) (*m.UserToken, error) {
	stmt, err := s.Prepare(q.GetUser())
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	var user m.User
	var password string
	err = stmt.QueryRow(u.Email).Scan(&user.Id, &user.Name, &user.Email, &password)
	if err != nil {
		return &m.UserToken{}, errors.New("Este usuario no existe.")
	}

	hashByte := []byte(password)
	passwordByte := []byte(u.Password)
	err = bcrypt.CompareHashAndPassword(hashByte, passwordByte) // Validación de la contrasenia por el hash
	if err != nil {
		return &m.UserToken{}, errors.New("Credenciales incorrectas.")
	}

	roles, _ := s.getRolesUser(user.Id)
	_token := s.srvJWT.GenerateToken(user, *roles)
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
	roles, _ := s.getRolesUser(user.Id)
	_token := s.srvJWT.GenerateToken(*user, *roles)
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
	err = stmt.QueryRow(id).Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		panic(err.Error())
	}

	return &user, nil
}

func (s *AuthService) getRolesUser(id int64) (*[]string, error) {
	stmt, err := s.Prepare(q.GetRolesUser())
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		panic(err.Error())
	}

	var rolesusers []string
	for rows.Next() {
		var rolesuser m.RolesUser
		err := rows.Scan(&rolesuser.Role)
		if err != nil {
			panic(err.Error())
		}

		rolesusers = append(rolesusers, rolesuser.Role)
	}

	if rows.Err() != nil {
		panic(err.Error())
	}

	return &rolesusers, nil
}

func (s *AuthService) createRoleUser(role string, id int64) (*int64, error) {
	stmt, err := s.Prepare(q.GetRoleId())
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	var roles_id int64
	_ = stmt.QueryRow(role).Scan(&roles_id)

	stmt1, err := s.Prepare(q.InsertRoleUser())
	if err != nil {
		panic(err.Error())
	}

	result, err := stmt1.Exec(roles_id, id)
	if err != nil {
		panic(err.Error())
	}

	rolesuser_id, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	return &rolesuser_id, nil
}
