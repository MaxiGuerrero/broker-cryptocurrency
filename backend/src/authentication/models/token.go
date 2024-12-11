package models

// Struct that represent the token of one user.
type UserToken struct {
	UserId string `json:"userId" bson:"_id,omitempty"`
	Token  string `json:"token"`
}
