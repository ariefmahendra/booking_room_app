package usecase

import (
	"booking-room/model/dto"
	"booking-room/repository"
	"booking-room/shared/service"
)

type AuthUC interface {
	Login(request dto.AuthRequest) (dto.AuthResponse, error)
}

type AuthUCImpl struct {
	employeeRepository repository.EmployeeRepository
	jwtService         service.JwtService
}

func (a *AuthUCImpl) Login(request dto.AuthRequest) (dto.AuthResponse, error) {
	//TODO implement me
	panic("implement me")
}
