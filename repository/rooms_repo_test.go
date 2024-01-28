package repository

import (
	"booking-room/model"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateRoom(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewRoomRepository(db)

	payload := model.Room{
		CodeRoom:   "101",
		RoomType:   "Standard",
		Capacity:   2,
		Facilities: "WiFi",
	}

	expectedRoom := model.Room{
		Id:         "1",
		CodeRoom:   "101",
		RoomType:   "Standard",
		Capacity:   2,
		Facilities: "WiFi",
	}

	mock.ExpectQuery("INSERT INTO rooms").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))

	room, err := repo.CreateRoom(payload)

	assert.NoError(t, err)
	assert.Equal(t, expectedRoom, room)
}

func TestUpdateRoom(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewRoomRepository(db)

	payload := model.Room{
		Id:         "1",
		CodeRoom:   "101",
		RoomType:   "Standard",
		Capacity:   2,
		Facilities: "WiFi",
	}

	mock.ExpectQuery("UPDATE rooms").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))

	room, err := repo.UpdateRoom(payload)

	assert.NoError(t, err)
	assert.Equal(t, payload, room)
}

func TestGetRoom(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewRoomRepository(db)

	expectedRoom := model.Room{
		Id:         "1",
		CodeRoom:   "101",
		RoomType:   "Standard",
		Capacity:   2,
		Facilities: "WiFi",
	}

	rows := sqlmock.NewRows([]string{"id", "code_room", "room_type", "capacity", "facilities"}).
		AddRow(expectedRoom.Id, expectedRoom.CodeRoom, expectedRoom.RoomType, expectedRoom.Capacity, expectedRoom.Facilities)

	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	room, err := repo.GetRoom(expectedRoom.Id)

	assert.NoError(t, err)
	assert.Equal(t, expectedRoom, room)
}

func TestListRoom(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewRoomRepository(db)

	expectedRooms := []model.Room{
		{
			Id:         "6368ecf2-012e-42f4-b707-4482188f72e5",
			CodeRoom:   "R001",
			RoomType:   "Meeting Room",
			Capacity:   25,
			Facilities: "Projector, Whiteboard, Audio System, Video Conferencing",
		},
		{
			Id:         "bef48a4a-5994-4130-a50f-6dd897dceafd",
			CodeRoom:   "R002",
			RoomType:   "Conference Room",
			Capacity:   20,
			Facilities: "Audio System, Video Conferencing, Large Screen Display",
		},
	}

	rows := sqlmock.NewRows([]string{"id", "code_room", "room_type", "capacity", "facilities"}).
		AddRow(expectedRooms[0].Id, expectedRooms[0].CodeRoom, expectedRooms[0].RoomType, expectedRooms[0].Capacity, expectedRooms[0].Facilities).
		AddRow(expectedRooms[1].Id, expectedRooms[1].CodeRoom, expectedRooms[1].RoomType, expectedRooms[1].Capacity, expectedRooms[1].Facilities)

	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	mock.ExpectQuery("SELECT COUNT").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))

	rooms, _, err := repo.ListRoom(1, 10)

	assert.NoError(t, err)
	assert.Equal(t, expectedRooms, rooms)
}

func TestRoomRepository_ListRoom_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewRoomRepository(db)

	mock.ExpectQuery("SELECT").WillReturnError(errors.New("failed to fetch rooms"))

	rooms, _, err := repo.ListRoom(1, 10)

	assert.Error(t, err)
	assert.Nil(t, rooms)
}
