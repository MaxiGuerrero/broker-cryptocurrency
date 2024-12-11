package authentication

import (
	"log"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

var costAlgorithmicEncrypter, _ = strconv.Atoi(os.Getenv("COST_ALGORITHMIC_ENCRYPTER"))

// Responsable to implement the logical encrypter. It implement the encrypter interface.
type Encrypter struct{}

// Implementation to generate a hash.
func (e Encrypter) GenerateHash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), costAlgorithmicEncrypter)
	if err != nil {
		log.Panicf("Cannot encrypt it, %v", err.Error())
	}
	return string(hash)
}

// Implementation to compare a plain password with hashed password gotten from database.
func (e Encrypter) Compare(hashedPassword, password string) bool {
	equal := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return equal == nil
}
