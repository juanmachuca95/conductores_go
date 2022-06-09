package interactor

import (
	"github.com/juanmachuca95/conductores_go/domains/models"
	"github.com/juanmachuca95/conductores_go/usescases/presenter"
	"github.com/juanmachuca95/conductores_go/usescases/repository"
)

type DriverInteractor interface {
	GetDriversWithPagination(int64) ([]*models.Driver, error)
	GetDriversAvailable() ([]*models.Driver, error)
}

type driverInteractor struct {
	driverRepository repository.DriverRepository
	driverPresenter  presenter.DriverPresenter
	dbRepository     repository.DBRepository
	jwtRepository    repository.JwtRepository
}

func NewDriverInteractor(d repository.DriverRepository, p presenter.DriverPresenter, b repository.DBRepository, j repository.JwtRepository) DriverInteractor {
	return &driverInteractor{driverRepository: d, driverPresenter: p, dbRepository: b, jwtRepository: j}
}

func (d *driverInteractor) GetDriversWithPagination(pager int64) ([]*models.Driver, error) {
	drivers, err := d.driverRepository.GetDriversWithPagination(pager)
	if err != nil {
		return nil, err
	}

	return d.driverPresenter.DriverResponse(drivers), nil
}

func (d *driverInteractor) GetDriversAvailable() ([]*models.Driver, error) {
	drivers, err := d.driverRepository.GetDriversAvailable()
	if err != nil {
		return nil, err
	}

	return d.driverPresenter.DriverResponse(drivers), nil
}
