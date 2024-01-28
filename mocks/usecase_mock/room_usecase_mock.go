package usecase_mock

import (
    "booking-room/model/dto"
    "booking-room/shared/shared_model"
    //"booking-room/repository"
    "github.com/stretchr/testify/mock"
)

type RoomUsecaseMock struct {
    mock.Mock
}

// RegisterNewRoom mocks the RegisterNewRoom method.
func (m *RoomUsecaseMock) RegisterNewRoom(payload dto.RoomRequest) (dto.RoomResponse, error) {
    args := m.Called(payload)
    return args.Get(0).(dto.RoomResponse), args.Error(1)
}

// FindRoomById mocks the FindRoomById method.
func (m *RoomUsecaseMock) FindRoomById(id string) (dto.RoomResponse, error) {
    args := m.Called(id)
    return args.Get(0).(dto.RoomResponse), args.Error(1)
}

// FindAllRoom mocks the FindAllRoom method.
func (m *RoomUsecaseMock) FindAllRoom(page, size int) ([]dto.RoomResponse, shared_model.Paging, error) {
    args := m.Called(page, size)
    return args.Get(0).([]dto.RoomResponse), args.Get(1).(shared_model.Paging), args.Error(2)
}

// UpdateRoom mocks the UpdateRoom method.
func (m *RoomUsecaseMock) UpdateRoom(payload dto.RoomRequest) (dto.RoomResponse, error) {
    args := m.Called(payload)
    return args.Get(0).(dto.RoomResponse), args.Error(1)
}
