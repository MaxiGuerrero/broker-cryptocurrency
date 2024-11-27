package authentication

type AuthRepository struct{}

func (a AuthRepository) CreateUser(username, password, email string) error {
	return nil
}
