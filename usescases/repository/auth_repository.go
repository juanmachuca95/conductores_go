package repository

import "github.com/juanmachuca95/conductores_go/domains/models"

type AuthRepository interface {
	Login(*models.Login) (*models.User, error)
}
