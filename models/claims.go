package models

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Username string   `json:"username"`
	Role     UserRole `json:"user_role"`
	jwt.RegisteredClaims
}
