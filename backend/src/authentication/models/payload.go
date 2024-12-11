package models

import (
	"time"
)

// Struct that represent the payload of the JWT token.
type Payload struct {
	UserId    string    `json:"userId" bson:"_id,omitempty"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at" bson:"deleted_at,omitempty"`
	Role      string    `json:"role"`
}
