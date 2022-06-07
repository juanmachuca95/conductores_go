package conductores

import (
	"database/sql"
	"log"

	srvJWT "github.com/juanmachuca95/conductores_go/internal/service_jwt"
	auth "github.com/juanmachuca95/conductores_go/services/auth/gateway"
	m "github.com/juanmachuca95/conductores_go/services/conductores/models"
	q "github.com/juanmachuca95/conductores_go/services/conductores/querys"
	"golang.org/x/crypto/bcrypt"
)

type ConductorStorage interface {
	getConductores(page int) (*[]m.Conductor, error)
	getConductoresDisponibles() (*[]m.Conductor, error)
	createConductor(u *m.CreateConductor) (*m.ConductorOk, error)
}

type ServiceConductor struct {
	*sql.DB
	srvJWT         srvJWT.JWTService
	authentication auth.AuthGateway
}

func NewConductorStorageGateway(db *sql.DB) *ServiceConductor {
	return &ServiceConductor{db, srvJWT.JWTAuthService(), auth.NewAuthGateway(db)}
}

func (s *ServiceConductor) createConductor(u *m.CreateConductor) (*m.ConductorOk, error) {
	stmt, err := s.Prepare(q.InsertUser())
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	hashPass, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	var passwordHashed string = string(hashPass)
	result, err := stmt.Exec(u.Name, u.Email, passwordHashed)
	if err != nil {
		return &m.ConductorOk{}, err
	}

	users_id, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	stmt1, err := s.Prepare(q.InsertConductor())
	if err != nil {
		panic(err.Error())
	}
	defer stmt1.Close()

	_, err = stmt1.Exec(users_id, u.Matricula, u.Vehiculo)
	if err != nil {
		panic(err.Error())
	}

	user, _ := s.authentication.GetUser(users_id)
	_, _ = s.authentication.CreateRoleUser("conductor", users_id)
	roles, _ := s.authentication.GetRolesUser(users_id)
	_token := s.srvJWT.GenerateToken(*user, *roles)

	return &m.ConductorOk{
		Token: _token,
	}, nil
}

func (s *ServiceConductor) getConductores(page int) (*[]m.Conductor, error) {
	stmt, err := s.Prepare(q.GetConductores())
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	perpage := 5
	if page == 0 {
		page = 1
	}

	rows, err := stmt.Query(perpage, ((page - 1) * perpage))
	if err != nil {
		log.Fatal(err.Error())
	}

	var conductores []m.Conductor
	for rows.Next() {
		var conductor m.Conductor
		err = rows.Scan(&conductor.Id, &conductor.Name, &conductor.Email, &conductor.Users_id, &conductor.Matricula, &conductor.Vehiculo, &conductor.Created_at, &conductor.Updated_at)
		if err != nil {
			panic(err.Error())
		}

		conductores = append(conductores, conductor)
	}

	if rows.Err() != nil {
		panic(rows.Err())
	}

	return &conductores, nil
}

func (s *ServiceConductor) getConductoresDisponibles() (*[]m.Conductor, error) {
	stmt, err := s.Prepare(q.GetConductoresDisponibles())
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}

	var conductores []m.Conductor
	for rows.Next() {
		var conductor m.Conductor
		err := rows.Scan(&conductor.Id, &conductor.Name, &conductor.Email, &conductor.Users_id, &conductor.Matricula, &conductor.Vehiculo, &conductor.Created_at, &conductor.Updated_at)

		if err != nil {
			panic(err)
		}

		conductores = append(conductores, conductor)
	}

	if rows.Err() != nil {
		panic(rows.Err())
	}

	return &conductores, nil
}
