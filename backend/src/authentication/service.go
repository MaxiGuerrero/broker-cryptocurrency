package authentication

type ICreateUserService interface {
	CreateUser(username, password, email, role string) error
}
