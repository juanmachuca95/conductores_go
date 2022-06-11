package registry

import (
	"github.com/juanmachuca95/conductores_go/interface/controllers"
	"github.com/juanmachuca95/conductores_go/interface/presenter"
	pi "github.com/juanmachuca95/conductores_go/interface/presenter"
	ri "github.com/juanmachuca95/conductores_go/interface/repository"
	"github.com/juanmachuca95/conductores_go/usescases/interactor"
	cr "github.com/juanmachuca95/conductores_go/usescases/repository"
)

func (r *registry) NewAuthController() controllers.AuthController {
	return controllers.NewAuthController(r.NewAuthInteractor())
}

func (r *registry) NewAuthInteractor() interactor.AuthInteractor {
	return interactor.NewAuthInteractor(r.NewAuthRepository(), r.NewAuthPresenter(), ri.NewDBRepository(r.db), ri.NewJwtRepository(), r.NewRolesRepository())
}

func (r *registry) NewAuthPresenter() presenter.AuthPresenter {
	return pi.NewAuthPresenter()
}

func (r *registry) NewAuthRepository() cr.AuthRepository {
	return ri.NewAuthRepository(r.db)
}

func (r *registry) NewRolesRepository() cr.RolesRepository {
	return ri.NewRolesRepository(r.db)
}
