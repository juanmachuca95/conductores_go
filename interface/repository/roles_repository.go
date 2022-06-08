package repository

import (
	"database/sql"

	"github.com/juanmachuca95/conductores_go/domains/models"
)

type rolesRepository struct {
	db *sql.DB
}

func (r *rolesRepository) GetRolesByUser(users_id int64) ([]*models.Role, error) {
	stmt, err := r.db.Prepare("SELECT r.role FROM roles AS r INNER JOIN rolesusers AS ro ON r.id = ro.roles_id WHERE ro.users_id = ?;")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(users_id)
	if err != nil {
		return nil, err
	}

	var roles []*models.Role
	for rows.Next() {
		var role models.Role
		err = rows.Scan(&role.Role)

		roles = append(roles, &role)
	}

	if rows.Err() != nil {
		panic(err)
	}

	return roles, nil
}
