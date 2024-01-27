package usecase

import (
	"booking-room/mocks/repo_mock"
	"booking-room/model"
	"booking-room/model/dto"
	"booking-room/shared/common"
	"booking-room/shared/shared_model"
	"fmt"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

var now = time.Date(2024, 1, 27, 7, 12, 4, 692220000, time.Local)

var expectedEmployees = []model.EmployeeModel{
	{
		Id:        "1",
		Name:      "eko",
		Email:     "eko@mail.com",
		Password:  "eko123",
		Division:  "HR",
		Position:  "Lead",
		Role:      "ADMIN",
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: &now,
	},
	{
		Id:        "2",
		Name:      "Juan",
		Email:     "juan@mail.com",
		Password:  "juan123",
		Division:  "IT",
		Position:  "Developer",
		Role:      "EMPLOYEE",
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: &now,
	},
}

var expectedPaging = shared_model.Paging{
	Page:        1,
	RowsPerPage: 5,
	TotalPages:  1,
	TotalRows:   2,
}

type EmployeeUCSuite struct {
	suite.Suite
	erm *repo_mock.EmployeeMock
	euc EmployeeUC
}

func TestEmployeeUCSuite(t *testing.T) {
	suite.Run(t, new(EmployeeUCSuite))
}

func (e *EmployeeUCSuite) SetupTest() {
	e.erm = new(repo_mock.EmployeeMock)
	e.euc = NewEmployeeUC(e.erm)
}

func (e *EmployeeUCSuite) TestGetDeletedEmployees_success() {
	e.erm.On("GetDeletedEmployees", expectedPaging.Page, expectedPaging.RowsPerPage).Return(expectedEmployees, expectedPaging, nil)

	var expectedResponses []dto.EmployeeResponse
	for _, employee := range expectedEmployees {
		employeeRes := common.EmployeeModelToResponse(employee)
		expectedResponses = append(expectedResponses, employeeRes)
	}

	employees, paging, err := e.euc.GetDeletedEmployees(expectedPaging.Page, expectedPaging.RowsPerPage)

	e.Nil(err)
	e.Equal(expectedResponses, employees)
	e.Equal(expectedPaging, paging)
}

func (e *EmployeeUCSuite) TestGetDeletedEmployees_failure() {
	e.erm.On("GetDeletedEmployees", expectedPaging.Page, expectedPaging.RowsPerPage).Return([]model.EmployeeModel{}, shared_model.Paging{}, fmt.Errorf("error"))

	employees, paging, err := e.euc.GetDeletedEmployees(expectedPaging.Page, expectedPaging.RowsPerPage)

	e.NotNil(err)
	e.Empty(employees)
	e.Empty(paging)
}

func (e *EmployeeUCSuite) TestGetEmployeeById_success() {
	e.erm.On("GetEmployeeById", expectedEmployees[0].Id).Return(expectedEmployees[0], nil)

	employee, err := e.euc.GetEmployeeById(expectedEmployees[0].Id)

	e.Nil(err)
	e.Equal(common.EmployeeModelToResponse(expectedEmployees[0]), employee)
}

func (e *EmployeeUCSuite) TestGetEmployeeById_failure() {
	e.erm.On("GetEmployeeById", expectedEmployees[0].Id).Return(model.EmployeeModel{}, fmt.Errorf("error"))

	employee, err := e.euc.GetEmployeeById(expectedEmployees[0].Id)

	e.NotNil(err)
	e.Equal(dto.EmployeeResponse{}, employee)
}

func (e *EmployeeUCSuite) TestGetEmployeeByEmail_success() {
	e.erm.On("GetEmployeeByEmail", expectedEmployees[0].Email).Return(expectedEmployees[0], nil)

	employee, err := e.euc.GetEmployeeByEmail(expectedEmployees[0].Email)

	e.Nil(err)
	e.Equal(common.EmployeeModelToResponse(expectedEmployees[0]), employee)
}

func (e *EmployeeUCSuite) TestGetEmployeeByEmail_failure() {
	e.erm.On("GetEmployeeByEmail", expectedEmployees[0].Email).Return(model.EmployeeModel{}, fmt.Errorf("error"))

	employee, err := e.euc.GetEmployeeByEmail(expectedEmployees[0].Email)

	e.NotNil(err)
	e.Equal(dto.EmployeeResponse{}, employee)
}

func (e *EmployeeUCSuite) TestDeleteEmployeeById_success() {
	e.erm.On("DeleteEmployeeById", expectedEmployees[0].Id).Return(nil)

	err := e.euc.DeleteEmployeeById(expectedEmployees[0].Id)

	e.Nil(err)
}

func (e *EmployeeUCSuite) TestDeleteEmployeeById_failure() {
	e.erm.On("DeleteEmployeeById", expectedEmployees[0].Id).Return(fmt.Errorf("error"))

	err := e.euc.DeleteEmployeeById(expectedEmployees[0].Id)

	e.NotNil(err)
}

func (e *EmployeeUCSuite) TestGetEmployees_success() {
	e.erm.On("GetEmployees", expectedPaging.Page, expectedPaging.RowsPerPage).Return(expectedEmployees, expectedPaging, nil)

	var expectedResponses []dto.EmployeeResponse
	for _, employee := range expectedEmployees {
		employeeRes := common.EmployeeModelToResponse(employee)
		expectedResponses = append(expectedResponses, employeeRes)
	}

	employees, paging, err := e.euc.GetEmployees(expectedPaging.Page, expectedPaging.RowsPerPage)

	e.Nil(err)
	e.Equal(expectedResponses, employees)
	e.Equal(expectedPaging, paging)
}

func (e *EmployeeUCSuite) TestGetEmployees_failure() {
	e.erm.On("GetEmployees", expectedPaging.Page, expectedPaging.RowsPerPage).Return([]model.EmployeeModel{}, shared_model.Paging{}, fmt.Errorf("error"))

	employees, paging, err := e.euc.GetEmployees(expectedPaging.Page, expectedPaging.RowsPerPage)

	e.NotNil(err)
	e.Empty(employees)
	e.Equal(shared_model.Paging{}, paging)
}

func (e *EmployeeUCSuite) TestCreateEmployee_success() {
	e.erm.On("InsertEmployee", mock.Anything).Return(expectedEmployees[0], nil)

	employee, err := e.euc.CreteEmployee(expectedEmployees[0])

	e.Nil(err)
	e.Equal(common.EmployeeModelToResponse(expectedEmployees[0]), employee)
}

func (e *EmployeeUCSuite) TestCreateEmployee_failure() {
	e.erm.On("InsertEmployee", mock.Anything).Return(model.EmployeeModel{}, fmt.Errorf("error"))

	employee, err := e.euc.CreteEmployee(expectedEmployees[0])

	e.NotNil(err)
	e.Empty(employee)
}

func (e *EmployeeUCSuite) TestUpdateEmployee_success() {
	e.erm.On("GetEmployeeById", expectedEmployees[0].Id).Return(expectedEmployees[0], nil)
	e.erm.On("UpdateEmployee", mock.Anything).Return(expectedEmployees[0], nil)

	employee, err := e.euc.UpdateEmployee(expectedEmployees[0])

	e.Nil(err)
	e.Equal(common.EmployeeModelToResponse(expectedEmployees[0]), employee)
}

func (e *EmployeeUCSuite) TestUpdateEmployee_failure() {
	e.erm.On("GetEmployeeById", expectedEmployees[0].Id).Return(model.EmployeeModel{}, fmt.Errorf("error"))

	employee, err := e.euc.UpdateEmployee(expectedEmployees[0])

	e.NotNil(err)
	e.Equal(dto.EmployeeResponse{}, employee)
}
