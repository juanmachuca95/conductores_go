package repository

import "github.com/juanmachuca95/conductores_go/domains/models"

type RolesRepository interface {
	GetRolesByUser(users_id int64) ([]*models.Role, error)
}
