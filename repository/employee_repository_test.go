package repository

import (
	"booking-room/config"
	"booking-room/model"
	"booking-room/shared/shared_model"
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"regexp"
	"testing"
	"time"
)

var now = time.Now()

var expectedEmployee = []model.EmployeeModel{{
	Id:        "9876234",
	Name:      "eko",
	Email:     "eko@mail.com",
	Password:  "12345",
	Division:  "IT",
	Position:  "Manager",
	Role:      "admin",
	Contact:   "1234567890",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
	DeletedAt: &now,
},
	{
		Id:        "9876234",
		Name:      "juan",
		Email:     "juan@mail.com",
		Password:  "12345",
		Division:  "IT",
		Position:  "Manager",
		Role:      "admin",
		Contact:   "1234567890",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	},
}

var ExpectedPaging = shared_model.Paging{
	Page:        1,
	RowsPerPage: 10,
	TotalPages:  5,
	TotalRows:   50,
}

type EmployeeRepositorySuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
	er      EmployeeRepository
}

func TestEmployeeRepositorySuite(t *testing.T) {
	suite.Run(t, new(EmployeeRepositorySuite))
}

func (e *EmployeeRepositorySuite) SetupSuite() {
	e.mockDb, e.mockSql, _ = sqlmock.New()
	e.er = NewEmployeeRepository(e.mockDb)
}

func (e *EmployeeRepositorySuite) TestCreateEmployee_success() {
	e.mockSql.ExpectQuery(regexp.QuoteMeta(config.InsertEmployee)).WithArgs(
		expectedEmployee[0].Name,
		expectedEmployee[0].Email,
		expectedEmployee[0].Password,
		expectedEmployee[0].Division,
		expectedEmployee[0].Position,
		expectedEmployee[0].Role,
		expectedEmployee[0].Contact,
	).WillReturnRows(sqlmock.NewRows([]string{
		"id",
		"name",
		"email",
		"division",
		"position",
		"role",
		"contact",
		"created_at",
	}).AddRow(
		expectedEmployee[0].Id,
		expectedEmployee[0].Name,
		expectedEmployee[0].Email,
		expectedEmployee[0].Division,
		expectedEmployee[0].Position,
		expectedEmployee[0].Role,
		expectedEmployee[0].Contact,
		expectedEmployee[0].CreatedAt,
	))

	employee, err := e.er.InsertEmployee(expectedEmployee[0])

	e.Nil(err)
	e.Equal(expectedEmployee[0].Name, employee.Name)
	e.Equal(expectedEmployee[0].Email, employee.Email)
	e.Equal(expectedEmployee[0].Division, employee.Division)
	e.Equal(expectedEmployee[0].Position, employee.Position)
	e.Equal(expectedEmployee[0].Role, employee.Role)
	e.Equal(expectedEmployee[0].Contact, employee.Contact)
}

func (e *EmployeeRepositorySuite) TestCreateEmployee_failure() {
	e.mockSql.ExpectQuery(regexp.QuoteMeta(config.InsertEmployee)).WithArgs(
		expectedEmployee[0].Name,
		expectedEmployee[0].Email,
		expectedEmployee[0].Password,
		expectedEmployee[0].Division,
		expectedEmployee[0].Position,
		expectedEmployee[0].Role,
		expectedEmployee[0].Contact,
	).WillReturnError(fmt.Errorf("error"))

	employee, err := e.er.InsertEmployee(expectedEmployee[0])

	e.NotNil(err)
	e.Equal(model.EmployeeModel{}, employee)
}

func (e *EmployeeRepositorySuite) TestUpdateEmployee_success() {
	e.mockSql.ExpectQuery(regexp.QuoteMeta(config.UpdateEmployeeById)).WithArgs(
		expectedEmployee[0].Name,
		expectedEmployee[0].Email,
		expectedEmployee[0].Password,
		expectedEmployee[0].Division,
		expectedEmployee[0].Position,
		expectedEmployee[0].Role,
		expectedEmployee[0].Contact,
		expectedEmployee[0].Id,
	).WillReturnRows(sqlmock.NewRows([]string{
		"id",
		"name",
		"email",
		"division",
		"position",
		"role",
		"contact",
		"created_at",
		"updated_at",
	}).AddRow(
		expectedEmployee[0].Id,
		expectedEmployee[0].Name,
		expectedEmployee[0].Email,
		expectedEmployee[0].Division,
		expectedEmployee[0].Position,
		expectedEmployee[0].Role,
		expectedEmployee[0].Contact,
		expectedEmployee[0].CreatedAt,
		expectedEmployee[0].UpdatedAt,
	))

	employee, err := e.er.UpdateEmployee(expectedEmployee[0])

	e.Nil(err)
	e.Equal(expectedEmployee[0].Name, employee.Name)
	e.Equal(expectedEmployee[0].Email, employee.Email)
	e.Equal(expectedEmployee[0].Division, employee.Division)
	e.Equal(expectedEmployee[0].Position, employee.Position)
	e.Equal(expectedEmployee[0].Role, employee.Role)
	e.Equal(expectedEmployee[0].Contact, employee.Contact)
}

func (e *EmployeeRepositorySuite) TestUpdateEmployee_failure() {
	e.mockSql.ExpectQuery(regexp.QuoteMeta(config.UpdateEmployeeById)).WithArgs(
		expectedEmployee[0].Name,
		expectedEmployee[0].Email,
		expectedEmployee[0].Password,
		expectedEmployee[0].Division,
		expectedEmployee[0].Position,
		expectedEmployee[0].Role,
		expectedEmployee[0].Contact,
		expectedEmployee[0].Id,
	).WillReturnError(fmt.Errorf("error"))

	employee, err := e.er.UpdateEmployee(expectedEmployee[0])

	e.NotNil(err)
	e.Equal(model.EmployeeModel{}, employee)
}

func (e *EmployeeRepositorySuite) TestDeleteEmployeeById_success() {
	e.mockSql.ExpectExec(regexp.QuoteMeta(config.DeleteEmployeeById)).WithArgs(
		expectedEmployee[0].Id,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	err := e.er.DeleteEmployeeById(expectedEmployee[0].Id)

	e.NoError(err)
}

func (e *EmployeeRepositorySuite) TestDeleteEmployeeById_failure() {
	e.mockSql.ExpectExec(regexp.QuoteMeta(config.DeleteEmployeeById)).WithArgs(
		expectedEmployee[0].Id,
	).WillReturnError(fmt.Errorf("error"))

	err := e.er.DeleteEmployeeById(expectedEmployee[0].Id)

	e.Error(err)
}

func (e *EmployeeRepositorySuite) TestGetEmployeeById_success() {
	e.mockSql.ExpectQuery(regexp.QuoteMeta(config.GetEmployeeById)).WithArgs(
		expectedEmployee[0].Id,
	).WillReturnRows(sqlmock.NewRows([]string{
		"id",
		"name",
		"email",
		"division",
		"position",
		"role",
		"contact",
		"created_at",
		"updated_at",
		"deleted_at",
	}).AddRow(
		expectedEmployee[0].Id,
		expectedEmployee[0].Name,
		expectedEmployee[0].Email,
		expectedEmployee[0].Division,
		expectedEmployee[0].Position,
		expectedEmployee[0].Role,
		expectedEmployee[0].Contact,
		expectedEmployee[0].CreatedAt,
		expectedEmployee[0].UpdatedAt,
		expectedEmployee[0].DeletedAt,
	))

	employeeById, err := e.er.GetEmployeeById(expectedEmployee[0].Id)

	e.Nil(err)
	e.Equal(expectedEmployee[0].Id, employeeById.Id)
	e.Equal(expectedEmployee[0].Name, employeeById.Name)
	e.Equal(expectedEmployee[0].Email, employeeById.Email)
	e.Equal(expectedEmployee[0].Division, employeeById.Division)
	e.Equal(expectedEmployee[0].Position, employeeById.Position)
	e.Equal(expectedEmployee[0].Role, employeeById.Role)
	e.Equal(expectedEmployee[0].Contact, employeeById.Contact)
}

func (e *EmployeeRepositorySuite) TestGetEmployeeById_failure() {
	e.mockSql.ExpectQuery(regexp.QuoteMeta(config.GetEmployeeById)).WithArgs(
		expectedEmployee[0].Id,
	).WillReturnError(fmt.Errorf("error"))

	employeeById, err := e.er.GetEmployeeById(expectedEmployee[0].Id)

	e.Error(err)
	e.Equal(model.EmployeeModel{}, employeeById)
}

func (e *EmployeeRepositorySuite) TestGetEmployeeByEmail_success() {
	e.mockSql.ExpectQuery(regexp.QuoteMeta(config.GetEmployeeByEmail)).WithArgs(
		expectedEmployee[0].Email,
	).WillReturnRows(sqlmock.NewRows([]string{
		"id",
		"name",
		"email",
		"division",
		"division",
		"position",
		"role",
		"contact",
		"created_at",
		"updated_at",
		"deleted_at",
	}).AddRow(
		expectedEmployee[0].Id,
		expectedEmployee[0].Name,
		expectedEmployee[0].Email,
		expectedEmployee[0].Password,
		expectedEmployee[0].Division,
		expectedEmployee[0].Position,
		expectedEmployee[0].Role,
		expectedEmployee[0].Contact,
		expectedEmployee[0].CreatedAt,
		expectedEmployee[0].UpdatedAt,
		expectedEmployee[0].DeletedAt,
	))

	employee, err := e.er.GetEmployeeByEmail(expectedEmployee[0].Email)

	e.Nil(err)
	e.Equal(expectedEmployee[0].Id, employee.Id)
	e.Equal(expectedEmployee[0].Name, employee.Name)
	e.Equal(expectedEmployee[0].Email, employee.Email)
	e.Equal(expectedEmployee[0].Division, employee.Division)
	e.Equal(expectedEmployee[0].Position, employee.Position)
	e.Equal(expectedEmployee[0].Role, employee.Role)
	e.Equal(expectedEmployee[0].Contact, employee.Contact)
	e.NotEmpty(employee.CreatedAt)
	e.NotEmpty(employee.UpdatedAt)
}

func (e *EmployeeRepositorySuite) TestGetEmployeeByEmail_failure() {
	e.mockSql.ExpectQuery(regexp.QuoteMeta(config.GetEmployeeByEmail)).WithArgs(
		expectedEmployee[0].Email,
	).WillReturnError(fmt.Errorf("error"))

	employee, err := e.er.GetEmployeeByEmail(expectedEmployee[0].Email)

	e.Error(err)
	e.Equal(model.EmployeeModel{}, employee)
}

func (e *EmployeeRepositorySuite) TestGetAllEmployees_success() {
	e.mockSql.ExpectQuery(regexp.QuoteMeta(config.GetEmployees)).WillReturnRows(sqlmock.NewRows([]string{
		"id",
		"name",
		"email",
		"division",
		"position",
		"role",
		"contact",
		"created_at",
		"updated_at",
		"deleted_at",
	}).AddRow(
		expectedEmployee[0].Id,
		expectedEmployee[0].Name,
		expectedEmployee[0].Email,
		expectedEmployee[0].Division,
		expectedEmployee[0].Position,
		expectedEmployee[0].Role,
		expectedEmployee[0].Contact,
		expectedEmployee[0].CreatedAt,
		expectedEmployee[0].UpdatedAt,
		expectedEmployee[0].DeletedAt,
	))

	e.mockSql.ExpectQuery(regexp.QuoteMeta(config.PagingEmployeeActive)).WillReturnRows(sqlmock.NewRows([]string{"total"}).AddRow(ExpectedPaging.TotalPages))

	employees, paging, err := e.er.GetEmployees(ExpectedPaging.Page, ExpectedPaging.RowsPerPage)

	e.Nil(err)
	e.Equal(1, len(employees))
	e.Equal(expectedEmployee[0].Id, employees[0].Id)
	e.Equal(expectedEmployee[0].Name, employees[0].Name)
	e.Equal(expectedEmployee[0].Email, employees[0].Email)
	e.Equal(expectedEmployee[0].Division, employees[0].Division)
	e.Equal(expectedEmployee[0].Position, employees[0].Position)
	e.Equal(expectedEmployee[0].Role, employees[0].Role)
	e.Equal(expectedEmployee[0].Contact, employees[0].Contact)
	e.NotEmpty(employees[0].CreatedAt)
	e.NotEmpty(employees[0].UpdatedAt)

	e.Equal(ExpectedPaging.Page, paging.Page)
	e.Equal(ExpectedPaging.RowsPerPage, paging.RowsPerPage)
}

func (e *EmployeeRepositorySuite) TestGetAllEmployees_failure() {
	e.mockSql.ExpectQuery(regexp.QuoteMeta(config.GetEmployees)).WillReturnError(fmt.Errorf("error"))

	employees, paging, err := e.er.GetEmployees(ExpectedPaging.Page, ExpectedPaging.RowsPerPage)

	e.Error(err)
	e.Equal([]model.EmployeeModel(nil), employees)
	e.Equal(shared_model.Paging{}, paging)
}

func (e *EmployeeRepositorySuite) TestGetDeletedEmployees_success() {
	e.mockSql.ExpectQuery(regexp.QuoteMeta(config.GetDeletedEmployees)).WillReturnRows(sqlmock.NewRows([]string{
		"id",
		"name",
		"email",
		"division",
		"position",
		"role",
		"contact",
		"created_at",
		"updated_at",
		"deleted_at",
	}).AddRow(
		expectedEmployee[0].Id,
		expectedEmployee[0].Name,
		expectedEmployee[0].Email,
		expectedEmployee[0].Division,
		expectedEmployee[0].Position,
		expectedEmployee[0].Role,
		expectedEmployee[0].Contact,
		expectedEmployee[0].CreatedAt,
		expectedEmployee[0].UpdatedAt,
		expectedEmployee[0].DeletedAt,
	))

	e.mockSql.ExpectQuery(regexp.QuoteMeta(config.PagingEmployeeDeleted)).WillReturnRows(sqlmock.NewRows([]string{"total"}).AddRow(ExpectedPaging.TotalPages))

	employees, paging, err := e.er.GetDeletedEmployees(ExpectedPaging.Page, ExpectedPaging.RowsPerPage)

	e.Nil(err)
	e.Equal(expectedEmployee[0].Id, employees[0].Id)
	e.Equal(expectedEmployee[0].Name, employees[0].Name)
	e.Equal(expectedEmployee[0].Email, employees[0].Email)
	e.Equal(expectedEmployee[0].Division, employees[0].Division)
	e.Equal(expectedEmployee[0].Position, employees[0].Position)
	e.Equal(expectedEmployee[0].Role, employees[0].Role)
	e.Equal(expectedEmployee[0].Contact, employees[0].Contact)
	e.NotEmpty(employees[0].CreatedAt)
	e.NotEmpty(employees[0].UpdatedAt)
	e.NotEmpty(employees[0].DeletedAt)

	e.Equal(ExpectedPaging.Page, paging.Page)
	e.Equal(ExpectedPaging.RowsPerPage, paging.RowsPerPage)
}

func (e *EmployeeRepositorySuite) TestGetDeletedEmployees_failure() {
	e.mockSql.ExpectQuery(regexp.QuoteMeta(config.GetDeletedEmployees)).WillReturnError(fmt.Errorf("error"))

	e.mockSql.ExpectQuery(regexp.QuoteMeta(config.PagingEmployeeDeleted)).WillReturnRows(sqlmock.NewRows([]string{"total"}).AddRow(ExpectedPaging.TotalPages))

	employees, paging, err := e.er.GetDeletedEmployees(ExpectedPaging.Page, ExpectedPaging.RowsPerPage)

	e.Error(err)
	e.Equal([]model.EmployeeModel(nil), employees)
	e.Equal(shared_model.Paging{}, paging)
}

