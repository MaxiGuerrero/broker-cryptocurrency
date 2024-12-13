package interfaces

import (
	"backend/src/system"
)

// Interface to implement methods about JWT management.
type IJWTBuilder interface {
	BuildToken(payload *system.Payload) string
	ValidateToken(tokenString string) (system.Payload, error)
}
