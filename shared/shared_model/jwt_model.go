package shared_model

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	AuthorId string `json:"authorId"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}
