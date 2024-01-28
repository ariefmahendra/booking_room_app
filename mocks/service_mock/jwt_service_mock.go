package service_mock

import (
	"booking-room/model/dto"
	"booking-room/shared/shared_model"
	"github.com/stretchr/testify/mock"
)

type JwtServiceMock struct {
	mock.Mock
}

func (j *JwtServiceMock) GenerateToken(employeeId, role string) (dto.AuthResponse, error) {
	args := j.Called(employeeId, role)
	return args.Get(0).(dto.AuthResponse), args.Error(1)
}

func (j *JwtServiceMock) ValidateToken(token string) (*shared_model.CustomClaims, error) {
	args := j.Called(token)
	return args.Get(0).(*shared_model.CustomClaims), args.Error(1)
}
