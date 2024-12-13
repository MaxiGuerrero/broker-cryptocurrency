package system

import (
	"errors"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

// Struct that represent the payload of the JWT token.
type Payload struct {
	UserId   string `json:"userId" bson:"_id,omitempty"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

// Responsable to implement the JWT token.
type JWTBuilder struct {
	secret []byte
}

func NewJWTBuilder() *JWTBuilder {
	var JWTSecret = []byte(os.Getenv("JWT_SECRET"))
	return &JWTBuilder{
		secret: JWTSecret,
	}
}

const (
	userId   = "userId"
	username = "username"
	role     = "role"
)

// Build token of a exists user.
func (j JWTBuilder) BuildToken(payload *Payload) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		userId:   payload.UserId,
		username: payload.Username,
		role:     payload.Role,
	})
	tokenString, err := token.SignedString(j.secret)
	if err != nil {
		log.Panicf("Error to generate token: %v", err.Error())
	}
	return tokenString
}

// Validate if a token is correct.
func (j JWTBuilder) ValidateToken(tokenString string) (Payload, error) {
	payload := Payload{}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return j.secret, nil
	})
	if err != nil {
		return payload, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return payload, errors.New("cannot claim the payload")
	}
	payload.UserId = claims[userId].(string)
	payload.Username = claims[username].(string)
	payload.Role = claims[role].(string)
	return payload, nil
}
