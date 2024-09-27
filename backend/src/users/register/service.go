package register

type ICreateUserService interface {
	CreateUser(username, password, email, role string) error
}
