package presenter

type AuthPresenter interface {
	AuthResponse(string) string
}

type authPresenter struct{}

func NewAuthPresenter() AuthPresenter {
	return &authPresenter{}
}

func (a *authPresenter) AuthResponse(token string) string {
	return token
}
