package repository

import (
	"booking-room/config"
	"booking-room/model"
	"booking-room/shared/shared_model"
	"database/sql"
	"fmt"
	"math"
)

type EmployeeRepository interface {
	InsertEmployee(payload model.EmployeeModel) (model.EmployeeModel, error)
	UpdateEmployee(payload model.EmployeeModel) (model.EmployeeModel, error)
	DeleteEmployeeById(id string) error
	GetEmployeeById(id string) (model.EmployeeModel, error)
	GetEmployeeByEmail(email string) (model.EmployeeModel, error)
	GetEmployees(page, size int) ([]model.EmployeeModel, shared_model.Paging, error)
	GetDeletedEmployees(page, size int) ([]model.EmployeeModel, shared_model.Paging, error)
}

type EmployeeRepositoryImpl struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &EmployeeRepositoryImpl{db: db}
}

func (e *EmployeeRepositoryImpl) GetDeletedEmployees(page, size int) ([]model.EmployeeModel, shared_model.Paging, error) {
	offset := (page - 1) * size

	var employees []model.EmployeeModel
	rows, err := e.db.Query(config.GetDeletedEmployees, size, offset)
	if err != nil {
		return nil, shared_model.Paging{}, fmt.Errorf("GetDeletedEmployees.Repository : %v", err)
	}

	for rows.Next() {
		var employee model.EmployeeModel
		err := rows.Scan(&employee.Id, &employee.Name, &employee.Email, &employee.Division, &employee.Position, &employee.Role, &employee.Contact, &employee.CreatedAt, &employee.UpdatedAt, &employee.DeletedAt)

		if err != nil {
			return nil, shared_model.Paging{}, fmt.Errorf("GetDeletedEmployees.Repository : %v", err)
		}

		employees = append(employees, employee)
	}

	totalRows := 0
	err = e.db.QueryRow(config.PagingEmployee).Scan(&totalRows)
	if err != nil {
		return nil, shared_model.Paging{}, fmt.Errorf("GetDeletedEmployees.Repository : %v", err)
	}

	paging := shared_model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}

	return employees, paging, nil
}

func (e *EmployeeRepositoryImpl) InsertEmployee(payload model.EmployeeModel) (model.EmployeeModel, error) {

	var employee model.EmployeeModel
	err := e.db.QueryRow(config.InsertEmployee, payload.Name, payload.Email, payload.Password, payload.Division, payload.Position, payload.Role, payload.Contact).Scan(&employee.Id, &employee.Name, &employee.Email, &employee.Division, &employee.Position, &employee.Role, &employee.Contact, &employee.CreatedAt)

	if err != nil {
		return model.EmployeeModel{}, fmt.Errorf("InsertEmployee.Repository : %v", err)
	}

	return employee, nil
}

func (e *EmployeeRepositoryImpl) UpdateEmployee(payload model.EmployeeModel) (model.EmployeeModel, error) {
	var employee model.EmployeeModel

	err := e.db.QueryRow(config.UpdateEmployeeById, payload.Name, payload.Email, payload.Password, payload.Division, payload.Position, payload.Role, payload.Contact, payload.Id).Scan(&employee.Id, &employee.Name, &employee.Email, &employee.Division, &employee.Position, &employee.Role, &employee.Contact, &employee.CreatedAt, &employee.UpdatedAt)
	if err != nil {
		return model.EmployeeModel{}, fmt.Errorf("UpdateEmployee.Repository : %v", err)
	}

	return employee, nil
}

func (e *EmployeeRepositoryImpl) DeleteEmployeeById(id string) error {
	_, err := e.db.Exec(config.EmployeeById, id)
	if err != nil {
		return err
	}

	return nil
}

func (e *EmployeeRepositoryImpl) GetEmployeeById(id string) (model.EmployeeModel, error) {
	var employee model.EmployeeModel

	err := e.db.QueryRow(config.GetEmployeeById, id).Scan(&employee.Id, &employee.Name, &employee.Email, &employee.Division, &employee.Position, &employee.Role, &employee.Contact, &employee.CreatedAt, &employee.UpdatedAt, &employee.DeletedAt)

	if err != nil {
		return model.EmployeeModel{}, fmt.Errorf("GetEmployeeById.Repository : %v", err)
	}

	return employee, nil
}

func (e *EmployeeRepositoryImpl) GetEmployeeByEmail(email string) (model.EmployeeModel, error) {
	var employee model.EmployeeModel

	err := e.db.QueryRow(config.GetEmployeeByEmail, email).Scan(&employee.Id, &employee.Name, &employee.Email, &employee.Password, &employee.Division, &employee.Position, &employee.Role, &employee.Contact, &employee.CreatedAt, &employee.UpdatedAt, &employee.DeletedAt)

	if err != nil {
		return model.EmployeeModel{}, fmt.Errorf("GetEmployeeByEmail.Repository : %v", err)
	}

	return employee, nil
}

func (e *EmployeeRepositoryImpl) GetEmployees(page, size int) ([]model.EmployeeModel, shared_model.Paging, error) {
	var employees []model.EmployeeModel
	var paging shared_model.Paging

	offset := (page - 1) * size

	rows, err := e.db.Query(config.GetEmployees, size, offset)
	if err != nil {
		return nil, shared_model.Paging{}, fmt.Errorf("GetEmployees.Repository : %v", err)
	}

	for rows.Next() {
		var employee model.EmployeeModel

		if err := rows.Scan(&employee.Id, &employee.Name, &employee.Email, &employee.Division, &employee.Position, &employee.Role, &employee.Contact, &employee.CreatedAt, &employee.UpdatedAt, &employee.DeletedAt); err != nil {
			return nil, shared_model.Paging{}, fmt.Errorf("GetEmployees.Repository : %v", err)
		}

		employees = append(employees, employee)
	}

	totalRows := 0

	if err = e.db.QueryRow(config.PagingEmployee).Scan(&totalRows); err != nil {
		return nil, shared_model.Paging{}, fmt.Errorf("GetEmployees.Repository : %v", err)
	}

	paging = shared_model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}

	return employees, paging, nil
}
