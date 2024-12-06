package interfaces

import (
	"backend/src/authentication/models"
)

// Interface to implement methods about JWT management.
type IJWTBuilder interface {
	BuildToken(payload *models.Payload) string
	ValidateToken(tokenString string) (*models.Payload, error)
}
