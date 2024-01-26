package usecase

import (
	"booking-room/mocks/repo_mock"
	"booking-room/model"
	"booking-room/shared/shared_model"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

var now = time.Now()

var expectedEmployees = []model.EmployeeModel{
	{
		Id:        "1",
		Name:      "test",
		Email:     "test",
		Password:  "test",
		Division:  "test",
		Position:  "test",
		Role:      "test",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: &now,
	},
	{
		Id:        "2",
		Name:      "test",
		Email:     "test",
		Password:  "test",
		Division:  "test",
		Position:  "test",
		Role:      "test",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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

func (e *EmployeeUCSuite) TestGetDeletedEmployees() {
	e.erm.On("GetDeletedEmployees").Return(expectedEmployees, nil)
	res, paging, err := e.euc.GetDeletedEmployees(expectedPaging.Page, expectedPaging.RowsPerPage)

	e.Nil(err)
	e.Equal(expectedPaging.Page, paging.Page)
	e.Equal(expectedPaging.RowsPerPage, paging.RowsPerPage)
}
