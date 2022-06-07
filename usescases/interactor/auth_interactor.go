package interactor

import (
	"log"

	"github.com/juanmachuca95/conductores_go/domains/models"
	"github.com/juanmachuca95/conductores_go/usescases/presenter"
	"github.com/juanmachuca95/conductores_go/usescases/repository"
)

type AuthInteractor interface {
	Authentication(*models.Login) (string, error)
}

type authInteractor struct {
	authRepository repository.AuthRepository
	authPresenter  presenter.AuthPresenter
}

func NewAuthInteractor(
	r repository.AuthRepository,
	p presenter.AuthPresenter,
) AuthInteractor {
	return &authInteractor{authRepository: r, authPresenter: p}
}

func (a *authInteractor) Authentication(u *models.Login) (string, error) {
	user, err := a.authRepository.Login(u)
	log.Println("authenticaction - user: ", user)
	if err != nil {
		return a.authPresenter.AuthResponse()
	}

	return a.authPresenter.AuthResponse()
}
