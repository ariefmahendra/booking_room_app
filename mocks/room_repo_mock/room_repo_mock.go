package room_repo_mock

import (
    "booking-room/model"
    "booking-room/shared/shared_model"
    "github.com/stretchr/testify/mock"
)

// RoomRepositoryMock adalah implementasi palsu dari RoomRepository untuk keperluan pengujian.
type RoomRepositoryMock struct {
    mock.Mock
}

// CreateRoom mocks the CreateRoom method.
func (m *RoomRepositoryMock) CreateRoom(payload model.Room) (model.Room, error) {
    args := m.Called(payload)
    return args.Get(0).(model.Room), args.Error(1)
}

// GetRoom mocks the GetRoom method.
func (m *RoomRepositoryMock) GetRoom(id string) (model.Room, error) {
    args := m.Called(id)
    return args.Get(0).(model.Room), args.Error(1)
}

// UpdateRoom mocks the UpdateRoom method.
func (m *RoomRepositoryMock) UpdateRoom(payload model.Room) (model.Room, error) {
    args := m.Called(payload)
    return args.Get(0).(model.Room), args.Error(1)
}

// ListRoom mocks the ListRoom method.
func (m *RoomRepositoryMock) ListRoom(page, size int) ([]model.Room, shared_model.Paging, error) {
    args := m.Called(page, size)
    return args.Get(0).([]model.Room), args.Get(1).(shared_model.Paging), args.Error(2)
}
