package interfaces

type IAuthService interface {
	Register(username, password, email string) error
}
