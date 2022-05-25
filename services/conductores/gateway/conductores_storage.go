package conductores

import (
	"database/sql"
	"log"

	srvJWT "github.com/juanmachuca95/spaceguru/internal/service_jwt"
	m "github.com/juanmachuca95/spaceguru/services/conductores/models"
	q "github.com/juanmachuca95/spaceguru/services/conductores/querys"
)

type ConductorStorage interface {
	getConductores(page int) (*[]m.Conductor, error)
	getConductoresDisponibles() (*[]m.Conductor, error)
}

type ServiceConductor struct {
	*sql.DB
	srvJWT srvJWT.JWTService
}

func NewConductorStorageGateway(db *sql.DB) *ServiceConductor {
	return &ServiceConductor{db, srvJWT.JWTAuthService()}
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
	log.Println("LIMIT ", perpage, " OFFSET ", ((page - 1) * perpage))
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
		err := rows.Scan(&conductor.Id, &conductor.Name, &conductor.Users_id, &conductor.Matricula, &conductor.Vehiculo, &conductor.Created_at, &conductor.Updated_at)

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
