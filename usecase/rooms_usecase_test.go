package usecase

import (
	"booking-room/mocks/room_usecase_mock"
	"booking-room/model"
	"booking-room/model/dto"
	"booking-room/shared/shared_model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	//"time"
)

var expectedRoom = model.Room{
	Id:         "a3d8e4ef-2e85-4ea5-9509-795f256226c3",
	CodeRoom:   "Ruang Candradimuka",
	RoomType:   "Ruang Meeting",
	Capacity:   42,
	Facilities: "available",
}

var expectedRooms = []model.Room{
	{
		Id:         "a3d8e4ef-2e85-4ea5-9509-795f256226c3",
		CodeRoom:   "Ruang Candradimuka",
		RoomType:   "Ruang Meeting",
		Capacity:   42,
		Facilities: "available",
	},
	{
		Id:         "e058c04a-7a41-4299-b618-a15e300b3554",
		CodeRoom:   "Ruang Bratasena",
		RoomType:   "Ruang Konferensi",
		Capacity:   21,
		Facilities: "booked",
	},
}

var page int = 1
var size int = 5

var expectedPaging = shared_model.Paging{
	Page:        page,
	RowsPerPage: size,
	TotalRows:   2,
	TotalPages:  1,
}

type RoomUseCaseTestSuite struct {
	suite.Suite
	mockUsecase *room_usecase_mock.RoomUsecaseMock
	uc          RoomUseCase
}

func (suite *RoomUseCaseTestSuite) SetupTest() {
	suite.mockUsecase = new(room_usecase_mock.RoomUsecaseMock)
	suite.uc = NewRoomUseCase(suite.mockUsecase)
}

func (suite *RoomUseCaseTestSuite) TestFindAllRoom() {
	suite.mockUsecase.On("FindAllRoom", page, size).Return(expectedRooms, expectedPaging, nil)

	actual, paging, err := suite.uc.FindAllRoom(page, size)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedRooms[0].CodeRoom, actual[0].CodeRoom)
	assert.Equal(suite.T(), expectedPaging.Page, paging.Page)
}

func (suite *RoomUseCaseTestSuite) TestFindroomByID() {
	suite.mockUsecase.On("FindRoomById", expectedRoom.Id).Return(dto.RoomResponse{
		Id:         expectedRoom.Id,
		CodeRoom:   expectedRoom.CodeRoom,
		RoomType:   expectedRoom.RoomType,
		Facilities: expectedRoom.Facilities,
		Capacity:   expectedRoom.Capacity,
	}, nil)

	actual, err := suite.uc.FindRoomById(expectedRoom.Id)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedRoom.CodeRoom, actual.CodeRoom)
}

func (suite *RoomUseCaseTestSuite) TestRegisterNewRoom_Success() {
	suite.mockUsecase.On("RegisterNewRoom", expectedRoom).Return(dto.RoomResponse{
		Id:         expectedRoom.Id,
		CodeRoom:   expectedRoom.CodeRoom,
		RoomType:   expectedRoom.RoomType,
		Facilities: expectedRoom.Facilities,
		Capacity:   expectedRoom.Capacity,
	}, nil)

	actual, err := suite.uc.RegisterNewRoom(expectedRoom)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedRoom.CodeRoom, actual.CodeRoom)
}

func (suite *RoomUseCaseTestSuite) TestRegisterNewRoom_EmptyStatusSuccess() {
	expectedPayload := expectedRoom
	expectedPayload.Facilities = ""

	suite.mockUsecase.On("RegisterNewRoom", expectedPayload).Return(dto.RoomResponse{
		Id:         expectedPayload.Id,
		CodeRoom:   expectedPayload.CodeRoom,
		RoomType:   expectedPayload.RoomType,
		Facilities: expectedPayload.Facilities,
		Capacity:   expectedPayload.Capacity,
	}, nil)

	_, err := suite.uc.RegisterNewRoom(expectedRoom)
	expectedRoom.Facilities = "available"

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *RoomUseCaseTestSuite) TestRegisterNewRoom_Failure() {
	suite.mockUsecase.On("RegisterNewRoom", expectedRoom).Return(dto.RoomResponse{}, someError)

	_, err := suite.uc.RegisterNewRoom(expectedRoom)

	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *RoomUseCaseTestSuite) TestRegisterNewRoom_EmptyFieldFailure() {
	expectedRoom.CodeRoom = ""
	_, err := suite.uc.RegisterNewRoom(expectedRoom)
	expectedRoom.CodeRoom = "Ruang Candradimuka"

	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *RoomUseCaseTestSuite) TestUpdateRoom_Success() {
	suite.mockUsecase.On("UpdateRoom", expectedRoom).Return(dto.RoomResponse{
		Id:         expectedRoom.Id,
		CodeRoom:   expectedRoom.CodeRoom,
		RoomType:   expectedRoom.RoomType,
		Facilities: expectedRoom.Facilities,
		Capacity:   expectedRoom.Capacity,
	}, nil)

	actual, err := suite.uc.UpdateRoom(expectedRoom)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedRoom.CodeRoom, actual.CodeRoom)
}

func (suite *RoomUseCaseTestSuite) TestUpdateRoom_Failure() {
	suite.mockUsecase.On("UpdateRoom", expectedRoom).Return(dto.RoomResponse{}, someError)

	_, err := suite.uc.UpdateRoom(expectedRoom)

	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *RoomUseCaseTestSuite) TestUpdateRoom_EmptyFieldFailure() {
	expectedRoom.CodeRoom = ""
	_, err := suite.uc.UpdateRoom(expectedRoom)
	expectedRoom.CodeRoom = "Ruang Candradimuka"

	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *RoomUseCaseTestSuite) TestUpdateRoom_EmptyStatusSuccess() {
	expectedPayload := expectedRoom
	expectedPayload.Facilities = ""

	suite.mockUsecase.On("UpdateRoom", expectedPayload).Return(dto.RoomResponse{
		Id:         expectedPayload.Id,
		CodeRoom:   expectedPayload.CodeRoom,
		RoomType:   expectedPayload.RoomType,
		Facilities: expectedPayload.Facilities,
		Capacity:   expectedPayload.Capacity,
	}, nil)

	_, err := suite.uc.UpdateRoom(expectedRoom)
	expectedRoom.Facilities = "available"

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

const someError = Error("some error")

type Error string

func (e Error) Error() string {
	return string(e)
}

func TestRoomUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(RoomUseCaseTestSuite))
}
