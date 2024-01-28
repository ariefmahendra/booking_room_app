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
	"time"
)

var now = time.Date(2024, 1, 27, 7, 12, 4, 692220000, time.Local)

var expectedRooms = []model.Room{
	{
		Id:           "1",
		CodeRoom:     "101",
		RoomType:     "Standard",
		Capacity:     2,
		Facilities:   "Basic Facilities",
		CreatedAt:    now,
		UpdatedAt:    now,
		DeletedAt:    &now,
	},
	{
		Id:           "2",
		CodeRoom:     "201",
		RoomType:     "Deluxe",
		Capacity:     4,
		Facilities:   "Luxury Facilities",
		CreatedAt:    now,
		UpdatedAt:    now,
		DeletedAt:    &now,
	},
}

type RoomUCSuite struct {
	suite.Suite
	rrm   *repo_mock.RoomMock
	ruc   RoomUC
}

func TestRoomUCSuite(t *testing.T) {
	suite.Run(t, new(RoomUCSuite))
}

func (r *RoomUCSuite) SetupTest() {
	r.rrm = new(repo_mock.RoomMock)
	r.ruc = NewRoomUC(r.rrm)
}

func (r *RoomUCSuite) TestGetDeletedRooms_success() {
	r.rrm.On("GetDeletedRooms", expectedPaging.Page, expectedPaging.RowsPerPage).Return(expectedRooms, expectedPaging, nil)

	var expectedResponses []dto.RoomResponse
	for _, room := range expectedRooms {
		roomRes := common.RoomModelToResponse(room)
		expectedResponses = append(expectedResponses, roomRes)
	}

	rooms, paging, err := r.ruc.GetDeletedRooms(expectedPaging.Page, expectedPaging.RowsPerPage)

	r.Nil(err)
	r.Equal(expectedResponses, rooms)
	r.Equal(expectedPaging, paging)
}

func (r *RoomUCSuite) TestGetDeletedRooms_failure() {
	r.rrm.On("GetDeletedRooms", expectedPaging.Page, expectedPaging.RowsPerPage).Return([]model.RoomModel{}, shared_model.Paging{}, fmt.Errorf("error"))

	rooms, paging, err := r.ruc.GetDeletedRooms(expectedPaging.Page, expectedPaging.RowsPerPage)

	r.NotNil(err)
	r.Empty(rooms)
	r.Empty(paging)
}

func (r *RoomUCSuite) TestGetRoomById_success() {
	r.rrm.On("GetRoomById", expectedRooms[0].Id).Return(expectedRooms[0], nil)

	room, err := r.ruc.GetRoomById(expectedRooms[0].Id)

	r.Nil(err)
	r.Equal(common.RoomModelToResponse(expectedRooms[0]), room)
}

func (r *RoomUCSuite) TestGetRoomById_failure() {
	r.rrm.On("GetRoomById", expectedRooms[0].Id).Return(model.RoomModel{}, fmt.Errorf("error"))

	room, err := r.ruc.GetRoomById(expectedRooms[0].Id)

	r.NotNil(err)
	r.Equal(dto.RoomResponse{}, room)
}

func (r *RoomUCSuite) TestGetRooms_success() {
	r.rrm.On("GetRooms", expectedPaging.Page, expectedPaging.RowsPerPage).Return(expectedRooms, expectedPaging, nil)

	var expectedResponses []dto.RoomResponse
	for _, room := range expectedRooms {
		roomRes := common.RoomModelToResponse(room)
		expectedResponses = append(expectedResponses, roomRes)
	}

	rooms, paging, err := r.ruc.GetRooms(expectedPaging.Page, expectedPaging.RowsPerPage)

	r.Nil(err)
	r.Equal(expectedResponses, rooms)
	r.Equal(expectedPaging, paging)
}

func (r *RoomUCSuite) TestGetRooms_failure() {
	r.rrm.On("GetRooms", expectedPaging.Page, expectedPaging.RowsPerPage).Return([]model.RoomModel{}, shared_model.Paging{}, fmt.Errorf("error"))

	rooms, paging, err := r.ruc.GetRooms(expectedPaging.Page, expectedPaging.RowsPerPage)

	r.NotNil(err)
	r.Empty(rooms)
	r.Equal(shared_model.Paging{}, paging)
}

func (r *RoomUCSuite) TestCreateRoom_success() {
	r.rrm.On("InsertRoom", mock.Anything).Return(expectedRooms[0], nil)

	room, err := r.ruc.CreateRoom(expectedRooms[0])

	r.Nil(err)
	r.Equal(common.RoomModelToResponse(expectedRooms[0]), room)
}

func (r *RoomUCSuite) TestCreateRoom_failure() {
	r.rrm.On("InsertRoom", mock.Anything).Return(model.RoomModel{}, fmt.Errorf("error"))

	room, err := r.ruc.CreateRoom(expectedRooms[0])

	r.NotNil(err)
	r.Empty(room)
}

func (r *RoomUCSuite) TestUpdateRoom_success() {
	r.rrm.On("GetRoomById", expectedRooms[0].Id).Return(expectedRooms[0], nil)
	r.rrm.On("UpdateRoom", mock.Anything).Return(expectedRooms[0], nil)

	room, err := r.ruc.UpdateRoom(expectedRooms[0])

	r.Nil(err)
	r.Equal(common.RoomModelToResponse(expectedRooms[0]), room)
}

func (r *RoomUCSuite) TestUpdateRoom_failure() {
	r.rrm.On("GetRoomById", expectedRooms[0].Id).Return(model.RoomModel{}, fmt.Errorf("error"))

	room, err := r.ruc.UpdateRoom(expectedRooms[0])

	r.NotNil(err)
	r.Equal(dto.RoomResponse{}, room)
}
