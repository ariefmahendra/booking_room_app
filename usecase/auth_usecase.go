package usecase

import (
	"booking-room/model/dto"
	"booking-room/repository"
	"booking-room/shared/common"
	"booking-room/shared/service"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
)

type AuthUC interface {
	Login(request dto.AuthRequest) (dto.AuthResponse, error)
	Register(request dto.EmployeeCreateRequest) (dto.EmployeeResponse, error)
}

type AuthUCImpl struct {
	employeeRepository repository.EmployeeRepository
	jwtService         service.JwtService
}

func NewAuthUC(employeeRepository repository.EmployeeRepository, jwtService service.JwtService) AuthUC {
	return &AuthUCImpl{employeeRepository: employeeRepository, jwtService: jwtService}
}

func (a *AuthUCImpl) Register(request dto.EmployeeCreateRequest) (dto.EmployeeResponse, error) {
	user := common.RequestToEmployeeModel(request)

	user.Role = strings.ToUpper(user.Role)
	if user.Role != "ADMIN" {
		log.Println("failed register user because invalid role")
		return dto.EmployeeResponse{}, errors.New("invalid role")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return dto.EmployeeResponse{}, err
	}

	user.Password = string(password)

	employee, err := a.employeeRepository.InsertEmployee(user)
	if err != nil {
		return dto.EmployeeResponse{}, err
	}

	return common.EmployeeModelToResponse(employee), nil
}

func (a *AuthUCImpl) Login(request dto.AuthRequest) (dto.AuthResponse, error) {
	userByEmail, err := a.employeeRepository.GetEmployeeByEmail(request.Email)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userByEmail.Password), []byte(request.Password)); err != nil {
		fmt.Println("middleware : invalid password")
		return dto.AuthResponse{}, err
	}

	if request.Email != userByEmail.Email {
		fmt.Println("middleware : invalid email")
		return dto.AuthResponse{}, errors.New("invalid email or email")
	}

	fmt.Println("Auth usecase : ", userByEmail.Id)
	authResponse, err := a.jwtService.GenerateToken(userByEmail.Id, userByEmail.Email, userByEmail.Role)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	return authResponse, nil
}
