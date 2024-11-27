package interfaces

type IAuthRepository interface {
	CreateUser(username, password, email string) error
}
