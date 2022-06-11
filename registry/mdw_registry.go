package registry

import (
	ri "github.com/juanmachuca95/conductores_go/interface/repository"
	"github.com/juanmachuca95/conductores_go/usescases/repository"
)

func (r *registry) NewMiddlewares() repository.MiddlewareRepository {
	return ri.NewMiddlewareRepository(ri.NewJwtRepository())
}
