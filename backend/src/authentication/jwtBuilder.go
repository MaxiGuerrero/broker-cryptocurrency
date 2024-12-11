package authentication

import (
	"backend/src/authentication/models"
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

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

// Build token of a exists user.
func (j JWTBuilder) BuildToken(payload *models.Payload) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    payload.UserId,
		"username":  payload.Username,
		"createdAt": payload.CreatedAt,
		"updatedAt": payload.UpdatedAt,
		"deletedAt": payload.DeletedAt,
		"role":      payload.Role,
	})
	tokenString, err := token.SignedString(j.secret)
	if err != nil {
		log.Panicf("Error to generate token: %v", err.Error())
	}
	return tokenString
}

// Validate if a token is correct.
func (j JWTBuilder) ValidateToken(tokenString string) (*models.Payload, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return j.secret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, errors.New("")
	}
	payload := &models.Payload{
		UserId:    claims["user_id"].(string),
		Username:  claims["username"].(string),
		CreatedAt: claims["created_at"].(time.Time),
		UpdatedAt: claims["updated_at"].(time.Time),
		DeletedAt: claims["deleted_at"].(time.Time),
		Role:      claims["role"].(string),
	}
	return payload, nil
}
