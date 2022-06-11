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
	authRepository  repository.AuthRepository
	authPresenter   presenter.AuthPresenter
	dbRepository    repository.DBRepository
	jwtRepository   repository.JwtRepository
	rolesRepository repository.RolesRepository
}

func NewAuthInteractor(
	a repository.AuthRepository,
	p presenter.AuthPresenter,
	d repository.DBRepository,
	j repository.JwtRepository,
	r repository.RolesRepository,
) AuthInteractor {
	return &authInteractor{
		authRepository:  a,
		authPresenter:   p,
		dbRepository:    d,
		jwtRepository:   j,
		rolesRepository: r,
	}
}

func (a *authInteractor) Authentication(u *models.Login) (string, error) {
	user, err := a.authRepository.Login(u)
	if err != nil {
		return "", err
	}

	roles, err := a.rolesRepository.GetRolesByUser(user)
	if err != nil {
		return "", err
	}

	token, err := a.jwtRepository.GenerateToken(user, roles)
	if err != nil {
		return "", err
	}

	return a.authPresenter.AuthResponse(token), nil
}
