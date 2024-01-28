package repository

import (
	"booking-room/config"
	"booking-room/model"
	"booking-room/model/dto"
	"booking-room/shared/shared_model"
	"database/sql"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

var newTime = time.Now()
var nowTest = newTime.Format("2006-01-02 15:04:05")
var expectedPaginf = shared_model.Paging{
	Page:        1,
	RowsPerPage: 5,
	TotalPages:  2,
	TotalRows:   10,
}

var expectedFacilities = model.Facilities{
	Id:             "1",
	CodeName:       "PRJ5",
	FacilitiesType: "projector",
	Status:         "AVAILABLE",
	CreatedAt:      nowTest,
	UpdatedAt:      nowTest,
	DeletedAt:      "",
}

var expectedFacilitiesResponse = []dto.FacilitiesResponse{
	{
		CodeName:       "PRJ5",
		FacilitiesType: "projector",
		Status:         "AVAILABLE",
	},
	{
		CodeName:       "PRJ6",
		FacilitiesType: "projector",
		Status:         "AVAILABLE",
	},
}

var expectedFacilitiesCreated = dto.FacilitiesCreated{
	Id:             "1",
	CodeName:       "PRJ5",
	FacilitiesType: "projector",
	Status:         "AVAILABLE",
	CreatedAt:      nowTest,
}
var inputUpdate = model.Facilities{
	CodeName:       "PRJ5",
	FacilitiesType: "projector",
	Status:         "AVAILABLE",
}
var expectedUpdate = dto.FacilitiesUpdated{
	Id:             "1",
	CodeName:       "PRJ5",
	FacilitiesType: "projector",
	Status:         "AVAILABLE",
	UpdatedAt:      nowTest,
}

type facilitiesRepositorySuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
	fr      FacilitiesRepository
}

func TestFacilitiesRepository(t *testing.T) {
	suite.Run(t, new(facilitiesRepositorySuite))
}

func (f *facilitiesRepositorySuite) SetupSuite() {
	f.mockDb, f.mockSql, _ = sqlmock.New()
	f.fr = NewFacilitiesRepository(f.mockDb)
}

func (f *facilitiesRepositorySuite) TestCreateFacilities_success() {
	f.mockSql.ExpectQuery(regexp.QuoteMeta(config.FacilityInsert)).WithArgs(
		expectedFacilities.CodeName,
		expectedFacilities.FacilitiesType,
	).WillReturnRows(sqlmock.NewRows([]string{
		"id",
		"status",
		"created_at",
	}).AddRow(
		expectedFacilities.Id,
		expectedFacilities.Status,
		expectedFacilities.CreatedAt,
	))
	facility, err := f.fr.Create(expectedFacilities)

	f.Nil(err)
	f.Equal(expectedFacilities.Id, facility.Id)
	f.Equal(expectedFacilitiesCreated.CodeName, facility.CodeName)
	f.Equal(expectedFacilitiesCreated.FacilitiesType, facility.FacilitiesType)
	f.Equal(expectedFacilitiesCreated.Status, facility.Status)
	f.Equal(expectedFacilitiesCreated.CreatedAt, facility.CreatedAt)
}

func (f *facilitiesRepositorySuite) TestCreateFacilities_failure() {
	f.mockSql.ExpectQuery(regexp.QuoteMeta(config.FacilityInsert)).WithArgs(
		expectedFacilities.CodeName,
		expectedFacilities.FacilitiesType,
	).WillReturnError(fmt.Errorf("error"))
	facility, err := f.fr.Create(expectedFacilities)

	f.NotNil(err)
	f.Equal(dto.FacilitiesCreated{}, facility)
}

func (f *facilitiesRepositorySuite) TestUpdateFacilities_success() {
	f.mockSql.ExpectQuery(regexp.QuoteMeta(config.FacilityUpdate)).WithArgs(
		inputUpdate.CodeName,
		inputUpdate.FacilitiesType,
		inputUpdate.Status,
		"1",
	).WillReturnRows(sqlmock.NewRows([]string{
		"id",
		"code_name",
		"facilities_type",
		"status",
		"updated_at",
	}).AddRow(
		expectedUpdate.Id,
		expectedUpdate.CodeName,
		expectedUpdate.FacilitiesType,
		expectedUpdate.Status,
		expectedUpdate.UpdatedAt,
	))
	facility, err := f.fr.Update(inputUpdate, "1")

	f.Nil(err)
	f.Equal(expectedUpdate.Id, facility.Id)
	f.Equal(expectedUpdate.CodeName, facility.CodeName)
	f.Equal(expectedUpdate.FacilitiesType, facility.FacilitiesType)
	f.Equal(expectedUpdate.Status, facility.Status)
	f.Equal(expectedUpdate.UpdatedAt, facility.UpdatedAt)
}

func (f *facilitiesRepositorySuite) TestUpdateFacilities_failure() {
	f.mockSql.ExpectQuery(regexp.QuoteMeta(config.FacilityUpdate)).WithArgs(
		inputUpdate.CodeName,
		inputUpdate.FacilitiesType,
		inputUpdate.Status,
		"1",
	).WillReturnError(fmt.Errorf("error"))
	facility, err := f.fr.Update(inputUpdate, "1")

	f.NotNil(err)
	f.Equal(dto.FacilitiesUpdated{}, facility)
}

func (f *facilitiesRepositorySuite) TestDeleteFacilities_success() {
	f.mockSql.ExpectExec(regexp.QuoteMeta(config.FacilityDeleteById)).WithArgs(
		"1",
	).WillReturnResult(sqlmock.NewResult(1, 1))
	err := f.fr.Delete("1")

	f.NoError(err)
}

/* func (f *facilitiesRepositorySuite) TestDeleteFacilities_failure() {
	f.mockSql.ExpectExec(regexp.QuoteMeta(config.FacilityDeleteById)).WithArgs(
		"1",
	).WillReturnError(fmt.Errorf("error"))
	err := f.fr.Delete("1")
	f.Error(err)
} */

func (f *facilitiesRepositorySuite) TestGetFacilityId() {
	f.mockSql.ExpectQuery(regexp.QuoteMeta(config.FacilityGetId)).WithArgs(
		"1",
	).WillReturnRows(sqlmock.NewRows([]string{
		"id",
		"code_name",
		"facilities_type",
		"status",
		"created_at",
		"updated_at",
		"deleted_at",
	}).AddRow(
		expectedFacilities.Id,
		expectedFacilities.CodeName,
		expectedFacilities.FacilitiesType,
		expectedFacilities.Status,
		expectedFacilities.CreatedAt,
		expectedFacilities.UpdatedAt,
		expectedFacilities.DeletedAt,
	))
	facility, err := f.fr.Get("1")

	f.Nil(err)
	f.Equal(expectedFacilities.Id, facility.Id)
	f.Equal(expectedFacilities.CodeName, facility.CodeName)
	f.Equal(expectedFacilities.FacilitiesType, facility.FacilitiesType)
	f.Equal(expectedFacilities.Status, facility.Status)
	f.Equal(expectedFacilities.CreatedAt, facility.CreatedAt)
	f.Equal(expectedFacilities.UpdatedAt, facility.UpdatedAt)
}

func (f *facilitiesRepositorySuite) TestGetFacilityId_failure() {
	f.mockSql.ExpectQuery(regexp.QuoteMeta(config.FacilityGetId)).WithArgs(
		"1",
	).WillReturnError(fmt.Errorf("error"))
	facility, err := f.fr.Get("1")
	f.Error(err)
	f.Equal(model.Facilities{}, facility)
}

func (f *facilitiesRepositorySuite) TestGetFacilityName_failure() {
	f.mockSql.ExpectQuery(regexp.QuoteMeta(config.FacilityGetName)).WithArgs(
		"1",
	).WillReturnError(fmt.Errorf("error"))
	facility, err := f.fr.GetName("1")

	f.NotNil(err)
	f.Equal(model.Facilities{}, facility)
}
func (f *facilitiesRepositorySuite) TestGetFacilityName() {
	f.mockSql.ExpectQuery(regexp.QuoteMeta(config.FacilityGetName)).WithArgs(
		"1",
	).WillReturnRows(sqlmock.NewRows([]string{
		"id",
		"code_name",
		"facilities_type",
		"status",
		"created_at",
		"updated_at",
		"deleted_at",
	}).AddRow(
		expectedFacilities.Id,
		expectedFacilities.CodeName,
		expectedFacilities.FacilitiesType,
		expectedFacilities.Status,
		expectedFacilities.CreatedAt,
		expectedFacilities.UpdatedAt,
		expectedFacilities.DeletedAt,
	))
	facility, err := f.fr.GetName("1")

	f.Nil(err)
	f.Equal(expectedFacilities.Id, facility.Id)
	f.Equal(expectedFacilities.CodeName, facility.CodeName)
	f.Equal(expectedFacilities.FacilitiesType, facility.FacilitiesType)
	f.Equal(expectedFacilities.Status, facility.Status)
	f.Equal(expectedFacilities.CreatedAt, facility.CreatedAt)
	f.Equal(expectedFacilities.UpdatedAt, facility.UpdatedAt)
}

/* func (f *facilitiesRepositorySuite) TestGetAllFacilities() {
	f.mockSql.ExpectQuery(regexp.QuoteMeta(config.RawPagingCount)).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(10))
	f.mockSql.ExpectQuery(regexp.QuoteMeta(config.FacilitiesList)).WithArgs(1, 5).
		WillReturnRows(sqlmock.NewRows([]string{
			"code_name",
			"facilities_type",
			"status",
		}).AddRow(
			expectedFacilitiesResponse[0].CodeName,
			expectedFacilitiesResponse[0].FacilitiesType,
			expectedFacilitiesResponse[0].Status,
		).AddRow(
			expectedFacilitiesResponse[1].CodeName,
			expectedFacilitiesResponse[1].FacilitiesType,
			expectedFacilitiesResponse[1].Status,
		))

	facilities, paging, err := f.fr.List(1, 5)

	f.Nil(err)
	f.Equal(expectedPaginf, paging)
	f.Equal(expectedFacilitiesResponse, facilities)
}
*/
