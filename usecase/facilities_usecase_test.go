package usecase

import (
	"booking-room/mocks/usecase_mock"
	"booking-room/model"
	"booking-room/model/dto"
	"booking-room/shared/shared_model"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var expectedFacilities = []dto.FacilitiesResponse{
	{CodeName: "SCR1", FacilitiesType: "screen", Status: "AVAILABLE"},
	{CodeName: "SCR2", FacilitiesType: "screen", Status: "BROKEN"},
}
var expectedFacility = model.Facilities{
	Id:             "randomuuid",
	CodeName:       "SCR1",
	FacilitiesType: "screen",
	Status:         "AVAILABLE",
	CreatedAt:      "2021-05-05",
	UpdatedAt:      "2021-05-05",
	DeletedAt:      "",
}
var expectedPagings = shared_model.Paging{
	Page:        1,
	TotalPages:  5,
	TotalRows:   10,
	RowsPerPage: 5,
}
var payloadCreate = model.Facilities{
	CodeName:       "SCR1",
	FacilitiesType: "screen",
}
var expectedCreate = dto.FacilitiesCreated{
	Id:             "randomuuid",
	CodeName:       "SCR1",
	FacilitiesType: "screen",
	Status:         "AVAILABLE",
	CreatedAt:      "2021-05-05",
}

type FacilitieUCSuite struct {
	suite.Suite
	mockUC       *usecase_mock.FacilitiesUseCaseMock
	facilitiesUC FacilitiesUsecase
}

func (suite *FacilitieUCSuite) SetupTest() {
	suite.mockUC = new(usecase_mock.FacilitiesUseCaseMock)
	suite.facilitiesUC = NewFacilitiesUsecase(suite.mockUC)
}

func TestFacilitiesUCSuite(t *testing.T) {
	suite.Run(t, new(FacilitieUCSuite))
}

func (suite *FacilitieUCSuite) TestList() {

	suite.mockUC.On("List", 1, 5).Return(expectedFacilities, expectedPaging, nil)

	facilities, paging, err := suite.facilitiesUC.List(1, 5)

	assert.NoError(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), expectedFacilities, facilities)
	assert.Equal(suite.T(), expectedPaging, paging)

	suite.mockUC.AssertExpectations(suite.T())
}

func (suite *FacilitieUCSuite) TestGetIdSuccess() {

	suite.mockUC.On("Get", "randomuuid").Return(expectedFacility, nil)

	facility, err := suite.facilitiesUC.Get("randomuuid")

	assert.NoError(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), expectedFacility, facility)

	suite.mockUC.AssertExpectations(suite.T())

}

func (suite *FacilitieUCSuite) TestGetIdFail() {
	suite.mockUC.On("Get", "randomuuid").Return(model.Facilities{}, errors.New("error"))

	facility, err := suite.facilitiesUC.Get("randomuuid")

	assert.Error(suite.T(), err, "Error should not be nil")
	assert.Equal(suite.T(), model.Facilities{}, facility)

	suite.mockUC.AssertExpectations(suite.T())
}

func (suite *FacilitieUCSuite) TestGetByCodeNameSuccess() {
	suite.mockUC.On("GetName", "SCR1").Return(expectedFacility, nil)

	facility, err := suite.facilitiesUC.GetByName("SCR1")

	assert.NoError(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), expectedFacility, facility)

	suite.mockUC.AssertExpectations(suite.T())
}

func (suite *FacilitieUCSuite) TestGetByCodeNameFail() {
	suite.mockUC.On("GetName", "Wrong Input").Return(model.Facilities{}, errors.New("error"))

	facility, err := suite.facilitiesUC.GetByName("Wrong Input")

	assert.Error(suite.T(), err, "Error should not be nil")
	assert.Equal(suite.T(), model.Facilities{}, facility)

	suite.mockUC.AssertExpectations(suite.T())
}

func (suite *FacilitieUCSuite) TestGetByTypeSuccess() {
	suite.mockUC.On("GetType", "screen", 1, 5).Return(expectedFacilities, expectedPaging, nil)

	facility, paging, err := suite.facilitiesUC.GetByType("screen", 1, 5)

	assert.NoError(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), expectedFacilities, facility)
	assert.Equal(suite.T(), expectedPaging, paging)

	suite.mockUC.AssertExpectations(suite.T())
}

func (suite *FacilitieUCSuite) TestGetByTypeFail() {
	suite.mockUC.On("GetType", "Wrong Input", 1, 5).Return([]dto.FacilitiesResponse{}, shared_model.Paging{}, errors.New("error"))

	facility, paging, err := suite.facilitiesUC.GetByType("Wrong Input", 1, 5)

	assert.Error(suite.T(), err, "Error should not be nil")
	assert.Equal(suite.T(), []dto.FacilitiesResponse{}, facility)
	assert.Equal(suite.T(), shared_model.Paging{}, paging)

	suite.mockUC.AssertExpectations(suite.T())
}

func (suite *FacilitieUCSuite) TestGetByStatusSuccess() {
	suite.mockUC.On("GetStatus", "AVAILABLE", 1, 5).Return(expectedFacilities, expectedPaging, nil)

	facility, paging, err := suite.facilitiesUC.GetByStatus("AVAILABLE", 1, 5)

	assert.NoError(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), expectedFacilities, facility)
	assert.Equal(suite.T(), expectedPaging, paging)

	suite.mockUC.AssertExpectations(suite.T())
}

func (suite *FacilitieUCSuite) TestGetByStatusFail() {
	suite.mockUC.On("GetStatus", "Wrong Input", 1, 5).Return([]dto.FacilitiesResponse{}, shared_model.Paging{}, errors.New("error"))

	facility, paging, err := suite.facilitiesUC.GetByStatus("Wrong Input", 1, 5)

	assert.Error(suite.T(), err, "Error should not be nil")
	assert.Equal(suite.T(), []dto.FacilitiesResponse{}, facility)
	assert.Equal(suite.T(), shared_model.Paging{}, paging)

	suite.mockUC.AssertExpectations(suite.T())
}

func (suite *FacilitieUCSuite) TestCreateSuccess() {
	suite.mockUC.On("Create", payloadCreate).Return(expectedCreate, nil)

	facility, err := suite.facilitiesUC.Create(payloadCreate)

	assert.NoError(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), expectedCreate, facility)

	suite.mockUC.AssertExpectations(suite.T())

}

func (suite *FacilitieUCSuite) TestDeletedSuccess() {
	suite.mockUC.On("Get", "randomuuid").Return(expectedFacility, nil)
	suite.mockUC.On("Delete", "randomuuid").Return(nil)
	expectedDel := errors.New("Facility deleted")

	err := suite.facilitiesUC.Delete("randomuuid")

	assert.Equal(suite.T(), err, expectedDel)

	suite.mockUC.AssertExpectations(suite.T())

}

func (suite *FacilitieUCSuite) TestDeletedByName() {
	suite.mockUC.On("GetName", "randomcodename").Return(expectedFacility, nil)
	suite.mockUC.On("DeleteByName", "randomcodename").Return(nil)
	expectedDel := errors.New("Facility deleted")

	err := suite.facilitiesUC.DeleteByName("randomcodename")

	assert.Equal(suite.T(), err, expectedDel)

	suite.mockUC.AssertExpectations(suite.T())

}

func (suite *FacilitieUCSuite) TestListGetDeleted() {
	expected := []model.Facilities{
		{
			Id:             "randomuuid",
			CodeName:       "SCR1",
			FacilitiesType: "screen",
			Status:         "AVAILABLE",
			CreatedAt:      "2021-05-05",
			UpdatedAt:      "2021-05-05",
			DeletedAt:      "",
		},
	}

	suite.mockUC.On("GetDeleted", 1, 5).Return(expected, expectedPaging, nil)

	facilities, paging, err := suite.facilitiesUC.GetDeleted(1, 5)

	assert.NoError(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), expected, facilities)
	assert.Equal(suite.T(), expectedPaging, paging)

	suite.mockUC.AssertExpectations(suite.T())
}

/* func (suite *FacilitieUCSuite) TestUpdateSuccess() {
	// Menyiapkan payload dan ID untuk tes
	payload := model.Facilities{
		CodeName:       "newCode",
		FacilitiesType: "newType",
		Status:         "newStatus",
	}
	id := "randomuid"

	suite.mockUC.On("Get", id).Return(expectedFacility, nil)

	suite.mockUC.
		On("Update", payload, id).Return(dto.FacilitiesUpdated{
		Id:             id,
		CodeName:       payload.CodeName,
		FacilitiesType: payload.FacilitiesType,
		Status:         payload.Status,
		UpdatedAt:      "now",
	}, nil)

	result, err := suite.mockUC.Update(payload, id)

	assert.NoError(suite.T(), err, "Error should be nil")

	expectedResult := dto.FacilitiesUpdated{
		Id:             id,
		CodeName:       payload.CodeName,
		FacilitiesType: payload.FacilitiesType,
		Status:         payload.Status,
		UpdatedAt:      "now",
	}
	assert.Equal(suite.T(), expectedResult, result, "Result should match")

	suite.mockUC.
		AssertExpectations(suite.T())
}
*/
