package conductores

import (
	"database/sql"

	m "github.com/juanmachuca95/spaceguru/services/conductores/models"
)

type ConductorGateway interface {
	GetConductores(page int) (*[]m.Conductor, error)
}

type ConductorInDB struct {
	ConductorStorage
}

func NewConductorGateway(db *sql.DB) ConductorGateway {
	return &ConductorInDB{NewConductorStorageGateway(db)}
}

// Obtener todos los conductores con paginación
func (c *ConductorInDB) GetConductores(page int) (*[]m.Conductor, error) {
	return c.getConductores(page)
}
