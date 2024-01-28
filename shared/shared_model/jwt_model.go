package shared_model

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}
