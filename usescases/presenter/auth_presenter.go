package presenter

type AuthPresenter interface {
	AuthResponse() (string, error)
}
