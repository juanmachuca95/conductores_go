package databases

import (
	"database/sql"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
)

type Seeder struct {
	*sql.DB
}

func NewSeeder(db *sql.DB) *Seeder { return &Seeder{db} }

func (s Seeder) UsersSeed() {
	stmt, _ := s.Prepare(`INSERT INTO users (name, email, password, created_at, updated_at) VALUES (?,?,?,NOW(),NOW())`)
	for i := 0; i < 100; i++ {
		_, err := stmt.Exec(faker.Name(), faker.Email(), faker.Password())
		if err != nil {
			panic(err)
		}
	}
}

func (s Seeder) RolesSeed() {
	rolesAllowed := []string{"conductor", "admin"}
	stmt, _ := s.Prepare("INSERT INTO roles (role, created_at, updated_at) VALUES (?,NOW(),NOW())")

	for _, role := range rolesAllowed {
		_, err := stmt.Exec(role)
		if err != nil {
			panic(err)
		}
	}
}

func (s Seeder) RolesUsersSeeder() {
	stmt, _ := s.Prepare("SELECT id, role FROM roles;")
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}

	var roles []struct {
		Id   int64
		Role string
	}
	for rows.Next() {
		var role struct {
			Id   int64
			Role string
		}
		err := rows.Scan(&role.Id, &role.Role)
		if err != nil {
			panic(err)
		}

		roles = append(roles, role)
	}

	stmt1, _ := s.Prepare("SELECT id FROM users;")
	defer stmt1.Close()

	stmt2, _ := s.Prepare("INSERT INTO rolesusers (roles_id, users_id, created_at, updated_at) VALUES (?,?, NOW(), NOW());")
	defer stmt2.Close()

	rowsUsers, err := stmt1.Query()
	if err != nil {
		panic(err)
	}
	rand.Seed(time.Now().Unix())
	for rowsUsers.Next() {
		var user struct {
			Id string
		}
		err := rowsUsers.Scan(&user.Id)
		if err != nil {
			panic(err)
		}

		n := rand.Int() % len(roles)
		_, err = stmt2.Exec(roles[n].Id, user.Id)
		if err != nil {
			panic(err)
		}
	}

}

func (s Seeder) ConductoresSeeder() {
	stmt, _ := s.Prepare("SELECT u.id FROM `users` as u INNER JOIN rolesusers as ru ON u.id = ru.users_id INNER JOIN roles as r ON r.id = ru.roles_id WHERE r.role = 'conductor';")
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}

	stmt1, _ := s.Prepare("INSERT INTO conductores (users_id, matricula, vehiculo, created_at, updated_at) VALUES (?,?,?,NOW(),NOW())")
	defer stmt1.Close()

	for rows.Next() {
		var users_id int64
		err := rows.Scan(&users_id)
		if err != nil {
			panic(err)
		}

		result, err := stmt1.Exec(users_id, rand.Intn(2000-0), faker.Name())
		if err != nil {
			panic(err.Error())
		}

		con_id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		s.ViajesSeeder(con_id)
	}
}

func (s Seeder) ViajesSeeder(con_id int64) {
	stmt, err := s.Prepare("INSERT INTO viajes (users_id, trip_status, created_at, updated_at) VALUES (?,?,NOW(),NOW())")
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	_, err = stmt.Exec(con_id, RandBool())
	if err != nil {
		panic(err.Error())
	}
}

func RandBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}
