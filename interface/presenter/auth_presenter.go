package presenter

type AuthPresenter interface {
	AuthResponse(string) (string, error)
}

type authPresenter struct{}

func NewAuthPresenter() AuthPresenter {
	return &authPresenter{}
}

func (a *authPresenter) AuthResponse(token string) (string, error) {
	return token, nil
}
