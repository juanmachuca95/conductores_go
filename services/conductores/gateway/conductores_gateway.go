package conductores

import (
	"database/sql"

	m "github.com/juanmachuca95/spaceguru/services/conductores/models"
)

type ConductorGateway interface {
	GetConductores(page int) (*[]m.Conductor, error)
	GetConductoresDisponibles() (*[]m.Conductor, error)
	CreateConductor(u *m.CreateConductor) (*m.ConductorOk, error)
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

// Obtener todos los conductores que no esten realizando viajes
func (c *ConductorInDB) GetConductoresDisponibles() (*[]m.Conductor, error) {
	return c.getConductoresDisponibles()
}

func (c *ConductorInDB) CreateConductor(u *m.CreateConductor) (*m.ConductorOk, error) {
	return c.createConductor(u)
}
