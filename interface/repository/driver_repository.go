package repository

import (
	"database/sql"
	"log"

	"github.com/juanmachuca95/conductores_go/domains/models"
	"github.com/juanmachuca95/conductores_go/usescases/repository"
)

type driverRepository struct {
	db *sql.DB
}

func NewDriverRepository(db *sql.DB) repository.DriverRepository {
	return &driverRepository{db: db}
}

func (d *driverRepository) GetDriversWithPagination(page int64) ([]*models.Driver, error) {
	stmt, err := d.db.Prepare("SELECT c.id, u.name, u.email, c.users_id, c.matricula, c.vehiculo, c.created_at, c.updated_at FROM conductores as c INNER JOIN users as u ON u.id = c.users_id LIMIT ? OFFSET ?;")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	var perpage int64 = 5
	if page == 0 {
		page = 1
	}

	rows, err := stmt.Query(perpage, ((page - 1) * perpage))
	if err != nil {
		log.Fatal(err.Error())
	}

	var drivers []*models.Driver
	for rows.Next() {
		var driver models.Driver
		err = rows.Scan(&driver.Id, &driver.Name, &driver.Email, &driver.Users_id, &driver.Matricula, &driver.Vehiculo, &driver.Created_at, &driver.Updated_at)
		if err != nil {
			panic(err.Error())
		}

		drivers = append(drivers, &driver)
	}

	if rows.Err() != nil {
		panic(rows.Err())
	}

	return drivers, nil
}

func (d *driverRepository) GetDriversAvailable() ([]*models.Driver, error) {
	stmt, err := d.db.Prepare("SELECT c.id, u.name, u.email, c.users_id, c.matricula, c.vehiculo, c.created_at, c.updated_at FROM conductores as c INNER JOIN users as u ON u.id = c.users_id INNER JOIN viajes v ON v.users_id = c.id WHERE trip_status=1;")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var drivers []*models.Driver
	for rows.Next() {
		var driver models.Driver
		err = rows.Scan(&driver.Id, &driver.Name, &driver.Email, &driver.Users_id, &driver.Matricula, &driver.Vehiculo, &driver.Created_at, &driver.Updated_at)

		drivers = append(drivers, &driver)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return drivers, nil
}
