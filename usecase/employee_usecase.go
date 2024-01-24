package usecase

import (
	"booking-room/model"
	"booking-room/model/dto"
	"booking-room/repository"
	"booking-room/shared/common"
	"booking-room/shared/shared_model"
	"fmt"
	"strings"
)

type EmployeeUC interface {
	CreteEmployee(payload model.EmployeeModel) (dto.EmployeeResponse, error)
	UpdateEmployee(payload model.EmployeeModel) (dto.EmployeeResponse, error)
	DeleteEmployeeById(id string) error
	GetEmployeeById(id string) (dto.EmployeeResponse, error)
	GetEmployeeByEmail(email string) (dto.EmployeeResponse, error)
	GetEmployees(page, size int) ([]dto.EmployeeResponse, shared_model.Paging, error)
}

type EmployeeUCImpl struct {
	employeeRepo repository.EmployeeRepository
}

func NewEmployeeUC(employeeRepo repository.EmployeeRepository) EmployeeUC {
	return &EmployeeUCImpl{employeeRepo: employeeRepo}
}

func (e *EmployeeUCImpl) CreteEmployee(payload model.EmployeeModel) (dto.EmployeeResponse, error) {
	payload.Role = strings.ToUpper(payload.Role)

	employee, err := e.employeeRepo.InsertEmployee(payload)
	if err != nil {
		return dto.EmployeeResponse{}, fmt.Errorf("CreateEmployee.Usecase : %v", err)
	}

	employeeDto := common.EmployeeModelToResponse(employee)

	return employeeDto, nil
}

func (e *EmployeeUCImpl) UpdateEmployee(payload model.EmployeeModel) (dto.EmployeeResponse, error) {
	payload.Role = strings.ToUpper(payload.Role)

	employee, err := e.employeeRepo.UpdateEmployee(payload)
	if err != nil {
		return dto.EmployeeResponse{}, fmt.Errorf("UpdateEmployee.Usecase : %v", err)
	}

	employeeDto := common.EmployeeModelToResponse(employee)

	return employeeDto, nil
}

func (e *EmployeeUCImpl) DeleteEmployeeById(id string) error {
	err := e.employeeRepo.DeleteEmployeeById(id)
	if err != nil {
		return fmt.Errorf("DeleteEmployeeById.Usecase : %v", err)
	}

	return nil
}

func (e *EmployeeUCImpl) GetEmployeeById(id string) (dto.EmployeeResponse, error) {
	employee, err := e.employeeRepo.GetEmployeeById(id)
	if err != nil {
		return dto.EmployeeResponse{}, fmt.Errorf("GetEmployeeById.Usecase : %v", err)
	}

	employeeDto := common.EmployeeModelToResponse(employee)

	return employeeDto, nil
}

func (e *EmployeeUCImpl) GetEmployeeByEmail(email string) (dto.EmployeeResponse, error) {
	employee, err := e.employeeRepo.GetEmployeeByEmail(email)
	if err != nil {
		return dto.EmployeeResponse{}, fmt.Errorf("GetEmployeeByEmail.Usecase : %v", err)
	}

	employeeDto := common.EmployeeModelToResponse(employee)

	return employeeDto, nil
}

func (e *EmployeeUCImpl) GetEmployees(page, size int) ([]dto.EmployeeResponse, shared_model.Paging, error) {
	// set default value for page and size
	if page == 0 && size == 0 {
		page, size = 1, 5
	}

	employees, paging, err := e.employeeRepo.GetEmployees(page, size)
	if err != nil {
		return nil, shared_model.Paging{}, fmt.Errorf("GetEmployees.Usecase : %v", err)
	}

	var employeesDto []dto.EmployeeResponse

	for _, employee := range employees {
		employeeDto := common.EmployeeModelToResponse(employee)
		employeesDto = append(employeesDto, employeeDto)
	}

	return employeesDto, paging, nil
}
