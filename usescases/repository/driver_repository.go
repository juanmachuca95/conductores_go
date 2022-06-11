package repository

import "github.com/juanmachuca95/conductores_go/domains/models"

type DriverRepository interface {
	GetDriversWithPagination(page int64) ([]*models.Driver, error)
	GetDriversAvailable() ([]*models.Driver, error)
}
