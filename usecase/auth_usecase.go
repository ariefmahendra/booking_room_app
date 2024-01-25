package usecase

import (
	"booking-room/model/dto"
	"booking-room/repository"
	"booking-room/shared/service"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type AuthUC interface {
	Login(request dto.AuthRequest) (dto.AuthResponse, error)
}

type AuthUCImpl struct {
	employeeRepository repository.EmployeeRepository
	jwtService         service.JwtService
}

func NewAuthUC(employeeRepository repository.EmployeeRepository, jwtService service.JwtService) AuthUC {
	return &AuthUCImpl{employeeRepository: employeeRepository, jwtService: jwtService}
}

func (a *AuthUCImpl) Login(request dto.AuthRequest) (dto.AuthResponse, error) {
	userByEmail, err := a.employeeRepository.GetEmployeeByEmail(request.Email)
	if err != nil {
		return dto.AuthResponse{}, err
	}
	fmt.Println(userByEmail)
	fmt.Println("userByEmail.password : ", userByEmail.Password)
	fmt.Println("request.Password : ", request.Password)
	if err := bcrypt.CompareHashAndPassword([]byte(userByEmail.Password), []byte(request.Password)); err != nil {
		fmt.Println("middleware : invalid password")
		return dto.AuthResponse{}, err
	}

	if request.Email != userByEmail.Email {
		fmt.Println("middleware : invalid email")
		return dto.AuthResponse{}, errors.New("invalid email or email")
	}

	authResponse, err := a.jwtService.GenerateToken(userByEmail.Id, userByEmail.Role)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	return authResponse, nil
}
