package usecase

import (
	"booking-room/mocks/repo_mock"
	"booking-room/model"
	"booking-room/model/dto"
	"booking-room/shared/common"
	"booking-room/shared/shared_model"
	"fmt"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

var expectedRoom = []model.Room{
	{
		Id:         "6368ecf2-012e-42f4-b707-4482188f72e5",
		CodeRoom:   "R001",
		RoomType:   "Meeting Room",
		Facilities: "Projector, Whiteboard, Audio System, Video Conferencing",
		Capacity:   25,
		CreatedAt:  now,
		UpdatedAt:  now,
		DeletedAt:  &now,
	},
}

type RoomUCSuite struct {
	suite.Suite
	roomRepoMock *repo_mock.RoomRepositoryMock
	roomUC       RoomUseCase
}

func TestRoomUCSuite(t *testing.T) {
	suite.Run(t, new(RoomUCSuite))
}

func (r *RoomUCSuite) SetupTest() {
	r.roomRepoMock = new(repo_mock.RoomRepositoryMock)
	r.roomUC = NewRoomUseCase(r.roomRepoMock)
}

func (r *RoomUCSuite) TestRegisterNewRoom_success() {
	r.roomRepoMock.On("CreateRoom", mock.Anything).Return(expectedRoom[0], nil)

	request := dto.RoomRequest{
		CodeRoom:   expectedRoom[0].CodeRoom,
		RoomType:   expectedRoom[0].RoomType,
		Facilities: expectedRoom[0].Facilities,
		Capacity:   expectedRoom[0].Capacity,
	}

	roomResponse, err := r.roomUC.RegisterNewRoom(request)

	r.Nil(err)
	r.Equal(common.RoomModelToResponse(expectedRoom[0]), roomResponse)
}

func (r *RoomUCSuite) TestRegisterNewRoom_failure() {
	r.roomRepoMock.On("CreateRoom", mock.Anything).Return(model.Room{}, fmt.Errorf("error"))

	request := dto.RoomRequest{
		CodeRoom:   expectedRoom[0].CodeRoom,
		RoomType:   expectedRoom[0].RoomType,
		Facilities: expectedRoom[0].Facilities,
		Capacity:   expectedRoom[0].Capacity,
	}

	roomResponse, err := r.roomUC.RegisterNewRoom(request)

	r.NotNil(err)
	r.Empty(roomResponse)
}

func (r *RoomUCSuite) TestUpdateRoom_success() {
	r.roomRepoMock.On("UpdateRoom", mock.Anything).Return(expectedRoom[0], nil)

	request := dto.RoomRequest{
		Id:         expectedRoom[0].Id,
		CodeRoom:   expectedRoom[0].CodeRoom,
		RoomType:   expectedRoom[0].RoomType,
		Facilities: expectedRoom[0].Facilities,
		Capacity:   expectedRoom[0].Capacity,
	}

	roomResponse, err := r.roomUC.UpdateRoom(request)

	r.Nil(err)
	r.Equal(common.RoomModelToResponse(expectedRoom[0]), roomResponse)
}

func (r *RoomUCSuite) TestUpdateRoom_failure() {
	r.roomRepoMock.On("UpdateRoom", mock.Anything).Return(model.Room{}, fmt.Errorf("error"))

	request := dto.RoomRequest{
		Id:         expectedRoom[0].Id,
		CodeRoom:   expectedRoom[0].CodeRoom,
		RoomType:   expectedRoom[0].RoomType,
		Facilities: expectedRoom[0].Facilities,
		Capacity:   expectedRoom[0].Capacity,
	}

	roomResponse, err := r.roomUC.UpdateRoom(request)

	r.NotNil(err)
	r.Equal(dto.RoomResponse{}, roomResponse)
}

func (r *RoomUCSuite) TestGetRoomById_success() {
	r.roomRepoMock.On("GetRoom", mock.Anything).Return(expectedRoom[0], nil)

	roomResponse, err := r.roomUC.FindRoomById(expectedRoom[0].Id)

	r.Nil(err)
	r.Equal(common.RoomModelToResponse(expectedRoom[0]), roomResponse)
}

func (r *RoomUCSuite) TestGetRoomById_failure() {
	r.roomRepoMock.On("GetRoom", mock.Anything).Return(model.Room{}, fmt.Errorf("error"))

	roomResponse, err := r.roomUC.FindRoomById(expectedRoom[0].Id)

	r.NotNil(err)
	r.Equal(dto.RoomResponse{}, roomResponse)
}

func (r *RoomUCSuite) TestGetListRoom_success() {
	r.roomRepoMock.On("ListRoom", expectedPaging.Page, expectedPaging.RowsPerPage).Return(expectedRoom, expectedPaging, nil)

	roomResponses, paging, err := r.roomUC.FindAllRoom(expectedPaging.Page, expectedPaging.RowsPerPage)

	var expectedResponse []dto.RoomResponse
	for _, room := range expectedRoom {
		roomResponse := common.RoomModelToResponse(room)
		expectedResponse = append(expectedResponse, roomResponse)
	}

	r.Nil(err)
	r.Equal(expectedResponse, roomResponses)
	r.Equal(expectedPaging, paging)
}

func (r *RoomUCSuite) TestGetListRoom_failure() {
	r.roomRepoMock.On("ListRoom", expectedPaging.Page, expectedPaging.RowsPerPage).Return([]model.Room{}, shared_model.Paging{}, fmt.Errorf("error"))

	roomResponses, paging, err := r.roomUC.FindAllRoom(expectedPaging.Page, expectedPaging.RowsPerPage)

	var expectedResponse []dto.RoomResponse
	for _, room := range expectedRoom {
		roomResponse := common.RoomModelToResponse(room)
		expectedResponse = append(expectedResponse, roomResponse)
	}

	r.NotNil(err)
	r.Empty(roomResponses)
	r.Empty(paging)
}
