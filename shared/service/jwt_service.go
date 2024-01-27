package service

import (
	"booking-room/config"
	"booking-room/model/dto"
	"booking-room/shared/shared_model"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtService interface {
	GenerateToken(id, email, role string) (dto.AuthResponse, error)
	ValidateToken(token string) (*shared_model.CustomClaims, error)
}

type JwtServiceImpl struct {
	cfg config.TokenConfig
}

func (j *JwtServiceImpl) GenerateToken(id, email, role string) (dto.AuthResponse, error) {
	claims := &shared_model.CustomClaims{
		Id:    id,
		Email: email,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.cfg.IssuerName,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.cfg.JwtExpiredTime)),
		},
	}

	token := jwt.NewWithClaims(j.cfg.JwtSigningMethod, claims)

	tokenSigned, err := token.SignedString(j.cfg.JwtSecretKey)
	if err != nil {
		return dto.AuthResponse{}, errors.New("error while signing token")
	}

	return dto.AuthResponse{Token: tokenSigned}, nil

}

func (j *JwtServiceImpl) ValidateToken(token string) (*shared_model.CustomClaims, error) {
	claims := &shared_model.CustomClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return j.cfg.JwtSecretKey, nil
	})
	if err != nil {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func NewJwtService(cfg config.TokenConfig) JwtService {
	return &JwtServiceImpl{cfg: cfg}
}
