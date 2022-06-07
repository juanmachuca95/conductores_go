package registry

import (
	"github.com/juanmachuca95/conductores_go/interface/controllers"
	pi "github.com/juanmachuca95/conductores_go/interface/presenter"
	ri "github.com/juanmachuca95/conductores_go/interface/repository"
	"github.com/juanmachuca95/conductores_go/usescases/presenter"
	"github.com/juanmachuca95/conductores_go/usescases/repository"
)

func (r *registry) NewAuthController() controllers.AuthController {
	return controllers.NewAuthController(r.NewAuthRepository(), r.NewAuthPresenter())
}

func (r *registry) NewAuthRepository() repository.AuthRepository {
	return ri.NewAuthRepository(r.db)
}

func (r *registry) NewAuthPresenter() presenter.AuthPresenter {
	return pi.NewAuthPresenter()
}
