package models

type Status int8

const (
	ACTIVE Status = iota
	BLOCKED
	INACTIVE
)

// Method to get the string of the status type.
func (s Status) String() string {
	return []string{"Active", "Blocked", "Inactive"}[s]
}
