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
var expectedPaging = shared_model.Paging{
	Page:        1,
	TotalPages:  5,
	TotalRows:   10,
	RowsPerPage: 5,
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
	suite.mockUC.On("GetName", "SCR1").Return(expectedFacilities, nil)

	facility, err := suite.facilitiesUC.GetByName("SCR1")

	assert.NoError(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), expectedFacilities, facility)

	suite.mockUC.AssertExpectations(suite.T())
}

func (suite *FacilitieUCSuite) TestGetByCodeNameFail() {
	suite.mockUC.On("GetName", "Wrong Input").Return(model.Facilities{}, errors.New("error"))

	facility, err := suite.facilitiesUC.GetByName("Wrong Input")

	assert.Error(suite.T(), err, "Error should not be nil")
	assert.Equal(suite.T(), model.Facilities{}, facility)

	suite.mockUC.AssertExpectations(suite.T())
}
