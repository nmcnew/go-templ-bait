package models

type UserRole int

const (
	Guest UserRole = iota
	User
	Moderator
	Admin
)

var userRoleName = map[UserRole]string{
	Guest:     "Guest",
	User:      "User",
	Moderator: "Moderator",
	Admin:     "Admin",
}

func (ur UserRole) String() string {
	return userRoleName[ur]
}

type ContextKeys int

const (
	ClaimsKey ContextKeys = iota
)
