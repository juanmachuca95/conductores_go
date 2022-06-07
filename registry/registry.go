package registry

import (
	"database/sql"

	"github.com/juanmachuca95/conductores_go/interface/controllers"
)

type Registry interface {
	NewAppController() controllers.AppController
}

type registry struct {
	db *sql.DB
}

func NewRegistry(db *sql.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controllers.AppController {
	return controllers.AppController{
		Auth: r.NewAuthController(),
	}
}
