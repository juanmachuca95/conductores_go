package presenter

import (
	"github.com/juanmachuca95/conductores_go/domains/models"
	"github.com/juanmachuca95/conductores_go/usescases/presenter"
)

type driverPresenter struct{}

func NewDriverPresenter() presenter.DriverPresenter {
	return &driverPresenter{}
}

func (d *driverPresenter) DriverResponse(drivers []*models.Driver) []*models.Driver {
	return drivers
}
