package presenter

type AuthPresenter interface {
	AuthResponse() (string, error)
}

type authPresenter struct{}

func NewAuthPresenter() AuthPresenter {
	return &authPresenter{}
}

func (a *authPresenter) AuthResponse() (string, error) {
	return "", nil
}
