package repo_mock

import (
	"booking-room/model"
	"booking-room/shared/shared_model"
	"github.com/stretchr/testify/mock"
)

type EmployeeMock struct {
	mock.Mock
}

func (e *EmployeeMock) GetDeletedEmployees(page, size int) ([]model.EmployeeModel, shared_model.Paging, error) {
	args := e.Called(page, size)
	return args.Get(0).([]model.EmployeeModel), args.Get(1).(shared_model.Paging), args.Error(2)
}

func (e *EmployeeMock) InsertEmployee(payload model.EmployeeModel) (model.EmployeeModel, error) {
	args := e.Called(payload)
	return args.Get(0).(model.EmployeeModel), args.Error(1)
}

func (e *EmployeeMock) UpdateEmployee(payload model.EmployeeModel) (model.EmployeeModel, error) {
	args := e.Called(payload)
	return args.Get(0).(model.EmployeeModel), args.Error(1)
}

func (e *EmployeeMock) DeleteEmployeeById(id string) error {
	args := e.Called(id)
	return args.Error(0)
}

func (e *EmployeeMock) GetEmployeeById(id string) (model.EmployeeModel, error) {
	args := e.Called(id)
	return args.Get(0).(model.EmployeeModel), args.Error(1)
}

func (e *EmployeeMock) GetEmployeeByEmail(email string) (model.EmployeeModel, error) {
	args := e.Called(email)
	return args.Get(0).(model.EmployeeModel), args.Error(1)
}

func (e *EmployeeMock) GetEmployees(page, size int) ([]model.EmployeeModel, shared_model.Paging, error) {
	args := e.Called(page, size)
	return args.Get(0).([]model.EmployeeModel), args.Get(1).(shared_model.Paging), args.Error(2)
}
