package usecase

import (
	"booking-room/model"
	"booking-room/model/dto"
	"booking-room/repository"
	"booking-room/shared/common"
	"booking-room/shared/shared_model"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type EmployeeUC interface {
	CreteEmployee(payload model.EmployeeModel) (dto.EmployeeResponse, error)
	UpdateEmployee(payload model.EmployeeModel) (dto.EmployeeResponse, error)
	DeleteEmployeeById(id string) error
	GetEmployeeById(id string) (dto.EmployeeResponse, error)
	GetEmployeeByEmail(email string) (dto.EmployeeResponse, error)
	GetEmployees(page, size int) ([]dto.EmployeeResponse, shared_model.Paging, error)
	GetDeletedEmployees(page, size int) ([]dto.EmployeeResponse, shared_model.Paging, error)
}

type EmployeeUCImpl struct {
	employeeRepo repository.EmployeeRepository
}

func NewEmployeeUC(employeeRepo repository.EmployeeRepository) EmployeeUC {
	return &EmployeeUCImpl{employeeRepo: employeeRepo}
}

func (e *EmployeeUCImpl) GetDeletedEmployees(page, size int) ([]dto.EmployeeResponse, shared_model.Paging, error) {
	if page == 0 && size == 0 {
		page, size = 1, 5
	}

	employees, paging, err := e.employeeRepo.GetDeletedEmployees(page, size)
	if err != nil {
		return nil, shared_model.Paging{}, err
	}

	var employeesDto []dto.EmployeeResponse
	for _, employee := range employees {
		employeeResponse := common.EmployeeModelToResponse(employee)
		employeesDto = append(employeesDto, employeeResponse)
	}

	return employeesDto, paging, nil
}

func (e *EmployeeUCImpl) CreteEmployee(payload model.EmployeeModel) (dto.EmployeeResponse, error) {
	payload.Role = strings.ToUpper(payload.Role)

	PasswordBytes, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 14)

	payload.Password = string(PasswordBytes)

	employee, err := e.employeeRepo.InsertEmployee(payload)
	if err != nil {
		return dto.EmployeeResponse{}, err
	}

	employeeDto := common.EmployeeModelToResponse(employee)

	return employeeDto, nil
}

func (e *EmployeeUCImpl) UpdateEmployee(payload model.EmployeeModel) (dto.EmployeeResponse, error) {

	_, err := e.employeeRepo.GetEmployeeById(payload.Id)
	if err != nil {
		return dto.EmployeeResponse{}, err
	}

	payload.Role = strings.ToUpper(payload.Role)

	employee, err := e.employeeRepo.UpdateEmployee(payload)
	if err != nil {
		return dto.EmployeeResponse{}, err
	}

	employeeDto := common.EmployeeModelToResponse(employee)

	return employeeDto, nil
}

func (e *EmployeeUCImpl) DeleteEmployeeById(id string) error {
	err := e.employeeRepo.DeleteEmployeeById(id)
	if err != nil {
		return err
	}

	return nil
}

func (e *EmployeeUCImpl) GetEmployeeById(id string) (dto.EmployeeResponse, error) {
	employee, err := e.employeeRepo.GetEmployeeById(id)
	if err != nil {
		return dto.EmployeeResponse{}, err
	}

	employeeDto := common.EmployeeModelToResponse(employee)

	return employeeDto, nil
}

func (e *EmployeeUCImpl) GetEmployeeByEmail(email string) (dto.EmployeeResponse, error) {
	employee, err := e.employeeRepo.GetEmployeeByEmail(email)
	if err != nil {
		return dto.EmployeeResponse{}, err
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
		return nil, shared_model.Paging{}, err
	}

	var employeesDto []dto.EmployeeResponse

	for _, employee := range employees {
		employeeDto := common.EmployeeModelToResponse(employee)
		employeesDto = append(employeesDto, employeeDto)
	}

	return employeesDto, paging, nil
}
