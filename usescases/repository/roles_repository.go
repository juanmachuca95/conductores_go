package repository

import "github.com/juanmachuca95/conductores_go/domains/models"

type RolesRepository interface {
	GetRolesByUser(user *models.User) ([]*models.Role, error)
}
