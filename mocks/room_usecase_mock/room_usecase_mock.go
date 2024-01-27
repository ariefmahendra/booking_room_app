package room_usecase_room

import (
	"booking-room/model"
	"booking-room/shared/shared_model"

	"github.com/stretchr/testify/mock"
)

type RoomUsecaseMock struct {
	mock.Mock
}

func (m *RoomUsecaseMock) RegisterNewRoom (payload model.Room) (model.Room, error) {
	args := m.Called(payload)
	return args.Get(0).(model.Room), args.Error(1)
}

func (m *RoomUsecaseMock) FindRoomById (Id string) (model.Room, error) {
	args := m.Called(Id)
	return args.Get(0).(model.Room), args.Error(1)
}

func (m *RoomUsecaseMock) FindAllRoom(page, size int) ([]model.Room, shared_model.Paging, error) {
	args := m.Called(page, size)
	return args.Get(0).([]model.Room), args.Get(1).(shared_model.Paging), args.Error(2)
}

func (m *RoomUsecaseMock) UpdateRoom(payload model.Room) (model.Room, error) {
	args := m.Called(payload)
	return args.Get(0).(model.Room), args.Error(1)
}

