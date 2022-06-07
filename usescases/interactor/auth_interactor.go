package interactor

import (
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
	dbRepository   repository.DBRepository
	jwtRepository  repository.JwtRepository
}

func NewAuthInteractor(
	r repository.AuthRepository,
	p presenter.AuthPresenter,
	d repository.DBRepository,
	j repository.JwtRepository,
) AuthInteractor {
	return &authInteractor{
		authRepository: r,
		authPresenter:  p,
		dbRepository:   d,
		jwtRepository:  j,
	}
}

func (a *authInteractor) Authentication(u *models.Login) (string, error) {
	user, err := a.authRepository.Login(u)
	if err != nil {
		return a.authPresenter.AuthResponse("")
	}

	token := a.jwtRepository.GenerateToken(*user, []string{"admin"})
	return a.authPresenter.AuthResponse(token)
}
