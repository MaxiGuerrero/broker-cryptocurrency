package models

import (
	"time"
)

type Status int8

const (
	Active Status = iota
	Blocked
	Inactive
)

// Method to get the string of the status type.
func (s Status) String() string {
	return []string{"Active", "Blocked", "Inactive"}[s]
}

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Role      string    `json:"role"`
}

type UserInfo struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
