package models

type Role int8

const (
	ADMIN Role = iota
	USER
)

// Method to get the string of the role type.
func (s Role) String() string {
	return []string{"ADMIN", "USER"}[s]
}
