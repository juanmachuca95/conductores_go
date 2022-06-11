package repository

import (
	"database/sql"

	"github.com/juanmachuca95/conductores_go/domains/models"
	"github.com/juanmachuca95/conductores_go/usescases/repository"
)

type rolesRepository struct {
	db *sql.DB
}

func NewRolesRepository(db *sql.DB) repository.RolesRepository {
	return &rolesRepository{db: db}
}

func (r *rolesRepository) GetRolesByUser(user *models.User) ([]*models.Role, error) {
	stmt, err := r.db.Prepare("SELECT r.role FROM roles AS r INNER JOIN rolesusers AS ro ON r.id = ro.roles_id WHERE ro.users_id = ?;")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(user.Id)
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
