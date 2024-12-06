package interfaces

// Interface to implement methods about hash management.
type IEncrypter interface {
	Compare(hashedPassword, password string) bool
	GenerateHash(password string) string
}
