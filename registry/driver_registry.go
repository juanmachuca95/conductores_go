package registry

import (
	"github.com/juanmachuca95/conductores_go/interface/controllers"
	pi "github.com/juanmachuca95/conductores_go/interface/presenter"
	ri "github.com/juanmachuca95/conductores_go/interface/repository"

	"github.com/juanmachuca95/conductores_go/usescases/interactor"
	"github.com/juanmachuca95/conductores_go/usescases/presenter"
	"github.com/juanmachuca95/conductores_go/usescases/repository"
)

func (r *registry) NewDriverController() controllers.DriverController {
	return controllers.NewDriverController(r.NewDriverInteractor())
}

func (r *registry) NewDriverInteractor() interactor.DriverInteractor {
	return interactor.NewDriverInteractor(r.NewDriverRepository(), r.NewDriverPresenter(), ri.NewDBRepository(r.db), ri.NewJwtRepository())
}

func (r *registry) NewDriverPresenter() presenter.DriverPresenter {
	return pi.NewDriverPresenter()
}

func (r *registry) NewDriverRepository() repository.DriverRepository {
	return ri.NewDriverRepository(r.db)
}
