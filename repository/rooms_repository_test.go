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
)

var expectedRoom = []model.Room{
	{
		Id:         "12345",
		CodeRoom:   "Room1",
		RoomType:   "Training",
		Facilities: "catering",
		Capacity:   25,
		CreatedAt:  now,
		UpdatedAt:  now,
		DeletedAt:  &now,
	},
}

type RoomsRepositorySuite struct {
	suite.Suite
	mockDb         *sql.DB
	mockSql        sqlmock.Sqlmock
	roomRepository RoomRepository
}

func TestRoomRepositorySuite(t *testing.T) {
	suite.Run(t, new(RoomsRepositorySuite))
}

func (r *RoomsRepositorySuite) SetupTest() {
	r.mockDb, r.mockSql, _ = sqlmock.New()
	r.roomRepository = NewRoomRepository(r.mockDb)
}

func (r *RoomsRepositorySuite) TestCreateRoom_success() {
	r.mockSql.ExpectQuery(regexp.QuoteMeta(config.CreateRoom)).WithArgs(expectedRoom[0].CodeRoom, expectedRoom[0].RoomType, expectedRoom[0].Capacity, expectedRoom[0].Facilities).WillReturnRows(sqlmock.NewRows([]string{
		"id",
		"code_room",
		"room_type",
		"facilities",
		"capacity",
		"created_at",
		"updated_at",
		"deleted_at",
	}).AddRow(
		expectedRoom[0].Id,
		expectedRoom[0].CodeRoom,
		expectedRoom[0].RoomType,
		expectedRoom[0].Facilities,
		expectedRoom[0].Capacity,
		expectedRoom[0].CreatedAt,
		expectedRoom[0].UpdatedAt,
		expectedRoom[0].DeletedAt,
	))

	room, err := r.roomRepository.CreateRoom(expectedRoom[0])

	r.Nil(err)
	r.Equal(expectedRoom[0], room)
}

func (r *RoomsRepositorySuite) TestCreateRoom_failure() {
	r.mockSql.ExpectQuery(regexp.QuoteMeta(config.CreateRoom)).WithArgs(expectedRoom[0].CodeRoom, expectedRoom[0].RoomType, expectedRoom[0].Capacity, expectedRoom[0].Facilities).WillReturnError(fmt.Errorf("error"))

	room, err := r.roomRepository.CreateRoom(expectedRoom[0])

	r.NotNil(err)
	r.Equal(model.Room{}, room)
}

func (r *RoomsRepositorySuite) TestUpdateRoom_success() {
	r.mockSql.ExpectQuery(regexp.QuoteMeta(config.UpdateRoomByID)).WithArgs(expectedRoom[0].Id, expectedRoom[0].CodeRoom, expectedRoom[0].RoomType, expectedRoom[0].Capacity, expectedRoom[0].Facilities).WillReturnRows(sqlmock.NewRows([]string{
		"id",
		"code_room",
		"room_type",
		"facilities",
		"capacity",
		"created_at",
		"updated_at",
		"deleted_at",
	}).AddRow(
		expectedRoom[0].Id,
		expectedRoom[0].CodeRoom,
		expectedRoom[0].RoomType,
		expectedRoom[0].Facilities,
		expectedRoom[0].Capacity,
		expectedRoom[0].CreatedAt,
		expectedRoom[0].UpdatedAt,
		expectedRoom[0].DeletedAt,
	))

	room, err := r.roomRepository.UpdateRoom(expectedRoom[0])

	r.Nil(err)
	r.Equal(expectedRoom[0], room)
}

func (r *RoomsRepositorySuite) TestUpdateRoom_failure() {
	r.mockSql.ExpectQuery(regexp.QuoteMeta(config.UpdateRoomByID)).WithArgs(expectedRoom[0].Id, expectedRoom[0].CodeRoom, expectedRoom[0].RoomType, expectedRoom[0].Capacity, expectedRoom[0].Facilities).WillReturnError(fmt.Errorf("error"))

	room, err := r.roomRepository.UpdateRoom(expectedRoom[0])

	r.NotNil(err)
	r.Equal(model.Room{}, room)
}

func (r *RoomsRepositorySuite) TestGetRoomById_success() {
	r.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectRoomByID)).WithArgs(expectedRoom[0].Id).WillReturnRows(sqlmock.NewRows([]string{
		"id",
		"code_room",
		"room_type",
		"capacity",
		"facilities",
		"created_at",
		"updated_at",
		"deleted_at",
	}).AddRow(
		expectedRoom[0].Id,
		expectedRoom[0].CodeRoom,
		expectedRoom[0].RoomType,
		expectedRoom[0].Capacity,
		expectedRoom[0].Facilities,
		expectedRoom[0].CreatedAt,
		expectedRoom[0].UpdatedAt,
		expectedRoom[0].DeletedAt,
	))

	room, err := r.roomRepository.GetRoom(expectedRoom[0].Id)

	r.Nil(err)
	r.Equal(expectedRoom[0], room)
}

func (r *RoomsRepositorySuite) TestGetRoomById_failure() {
	r.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectRoomByID)).WithArgs(expectedRoom[0].Id).WillReturnError(fmt.Errorf("error"))

	room, err := r.roomRepository.GetRoom(expectedRoom[0].Id)

	r.NotNil(err)
	r.Equal(model.Room{}, room)
}

func (r *RoomsRepositorySuite) TestGetListRoom_success() {
	offset := (ExpectedPaging.Page - 1) * ExpectedPaging.RowsPerPage

	r.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectRoomList)).WithArgs(ExpectedPaging.RowsPerPage, offset).WillReturnRows(sqlmock.NewRows([]string{
		"id",
		"code_room",
		"room_type",
		"capacity",
		"facilities",
		"created_at",
		"updated_at",
		"deleted_at",
	}).AddRow(
		expectedRoom[0].Id,
		expectedRoom[0].CodeRoom,
		expectedRoom[0].RoomType,
		expectedRoom[0].Capacity,
		expectedRoom[0].Facilities,
		expectedRoom[0].CreatedAt,
		expectedRoom[0].UpdatedAt,
		expectedRoom[0].DeletedAt,
	))

	r.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectCountRoom)).WillReturnRows(
		sqlmock.NewRows([]string{
			"total",
		}).AddRow(
			ExpectedPaging.TotalRows,
		))

	room, paging, err := r.roomRepository.ListRoom(ExpectedPaging.Page, ExpectedPaging.RowsPerPage)

	r.Nil(err)
	r.Equal(ExpectedPaging, paging)
	r.Equal(expectedRoom, room)
}

func (r *RoomsRepositorySuite) TestGetListRoom_failure() {
	offset := (ExpectedPaging.Page - 1) * ExpectedPaging.RowsPerPage

	r.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectRoomList)).WithArgs(ExpectedPaging.RowsPerPage, offset).WillReturnError(fmt.Errorf("error"))

	room, paging, err := r.roomRepository.ListRoom(ExpectedPaging.Page, ExpectedPaging.RowsPerPage)

	r.NotNil(err)
	r.Equal(shared_model.Paging{}, paging)
	r.Empty(room)
}
