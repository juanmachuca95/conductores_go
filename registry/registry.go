package registry

import (
	"database/sql"

	"github.com/juanmachuca95/conductores_go/interface/controllers"
	"github.com/juanmachuca95/conductores_go/usescases/repository"
)

type Registry interface {
	NewAppController() controllers.AppController
	NewMiddlewares() repository.MiddlewareRepository
}

type registry struct {
	db *sql.DB
}

func NewRegistry(db *sql.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controllers.AppController {
	return controllers.AppController{
		Auth:   r.NewAuthController(),
		Driver: r.NewDriverController(),
	}
}
