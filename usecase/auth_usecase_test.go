package usecase

import (
	"booking-room/mocks/repo_mock"
	"booking-room/mocks/service_mock"
	"booking-room/model"
	"booking-room/model/dto"
	"booking-room/shared/common"
	"fmt"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

type AuthUCSuite struct {
	suite.Suite
	jwtServiceMock *service_mock.JwtServiceMock
	erm            *repo_mock.EmployeeMock
	authUC         AuthUC
}

func (a *AuthUCSuite) SetupTest() {
	a.jwtServiceMock = new(service_mock.JwtServiceMock)
	a.erm = new(repo_mock.EmployeeMock)
	a.authUC = NewAuthUC(a.erm, a.jwtServiceMock)
}

func TestAuthUCSuite(t *testing.T) {
	suite.Run(t, new(AuthUCSuite))
}

func (a *AuthUCSuite) TestRegister_success() {
	a.erm.On("InsertEmployee", mock.Anything).Return(expectedEmployees[0], nil)

	employeeRequest := dto.EmployeeCreateRequest{
		Name:     expectedEmployees[0].Name,
		Email:    expectedEmployees[0].Email,
		Password: expectedEmployees[0].Password,
		Role:     expectedEmployees[0].Role,
		Division: expectedEmployees[0].Division,
		Position: expectedEmployees[0].Position,
		Contact:  expectedEmployees[0].Contact,
	}

	response, err := a.authUC.Register(employeeRequest)

	a.Nil(err)
	a.Equal(common.EmployeeModelToResponse(expectedEmployees[0]), response)
}

func (a *AuthUCSuite) TestRegister_failure() {
	a.erm.On("InsertEmployee", mock.Anything).Return(model.EmployeeModel{}, fmt.Errorf("failed insert employee"))

	employeeRequest := dto.EmployeeCreateRequest{
		Name:     expectedEmployees[0].Name,
		Email:    expectedEmployees[0].Email,
		Password: expectedEmployees[0].Password,
		Role:     expectedEmployees[0].Role,
		Division: expectedEmployees[0].Division,
		Position: expectedEmployees[0].Position,
		Contact:  expectedEmployees[0].Contact,
	}

	response, err := a.authUC.Register(employeeRequest)

	a.NotNil(err)
	a.Empty(response)
}

func (a *AuthUCSuite) TestLogin_success() {
	authRequest := dto.AuthRequest{
		Email:    expectedEmployees[0].Email,
		Password: expectedEmployees[0].Password,
	}

	authResponse := dto.AuthResponse{
		Token: "token",
	}

	password, err := bcrypt.GenerateFromPassword([]byte(expectedEmployees[0].Password), 14)
	a.Nil(err)

	expectedEmployees[0].Password = string(password)

	a.erm.On("GetEmployeeByEmail", authRequest.Email).Return(expectedEmployees[0], nil)

	a.jwtServiceMock.On("GenerateToken", expectedEmployees[0].Id, expectedEmployees[0].Email, expectedEmployees[0].Role).Return(authResponse, nil)

	response, err := a.authUC.Login(authRequest)

	a.Nil(err)
	a.Equal(authResponse, response)
}

func (a *AuthUCSuite) TestLogin_failure() {
	authRequest := dto.AuthRequest{
		Email:    expectedEmployees[0].Email,
		Password: expectedEmployees[0].Password,
	}

	a.erm.On("GetEmployeeByEmail", expectedEmployees[0].Email).Return(model.EmployeeModel{}, fmt.Errorf("failed get employee by email"))

	response, err := a.authUC.Login(authRequest)

	a.NotNil(err)
	a.Empty(response)
}
