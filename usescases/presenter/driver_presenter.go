package presenter

import "github.com/juanmachuca95/conductores_go/domains/models"

type DriverPresenter interface {
	DriverResponse(drivers []*models.Driver) []*models.Driver
}
