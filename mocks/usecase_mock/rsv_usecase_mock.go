package usecase_mock

import (
	"booking-room/model/dto"
	"booking-room/shared/shared_model"

	"github.com/stretchr/testify/mock"
)

type RsvUseCaseMock struct {
	mock.Mock
}

func (m *RsvUseCaseMock) List(page, size int) ([]dto.TransactionDTO, shared_model.Paging, error){
	args := m.Called(page, size)
	return args.Get(0).([]dto.TransactionDTO), args.Get(1).(shared_model.Paging), args.Error(2)
}
func (m *RsvUseCaseMock) GetID(id string) (dto.TransactionDTO, error){
	args := m.Called(id)
	return args.Get(0).(dto.TransactionDTO), args.Error(1)
}
func (m *RsvUseCaseMock) GetEmployee(id string, page, size int) ([]dto.TransactionDTO, shared_model.Paging, error){
	args := m.Called(id, page, size)
	return args.Get(0).([]dto.TransactionDTO), args.Get(1).(shared_model.Paging), args.Error(2)
}
func (m *RsvUseCaseMock) PostReservation(payload dto.PayloadReservationDTO) (dto.TransactionDTO, error){
	args := m.Called(payload)
	return args.Get(0).(dto.TransactionDTO), args.Error(1)
}
func (m *RsvUseCaseMock) UpdateStatus(payload dto.TransactionDTO) (dto.TransactionDTO, error){
	args := m.Called(payload)
	return args.Get(0).(dto.TransactionDTO), args.Error(1)
}
func (m *RsvUseCaseMock) DeleteResv(id string) (string, error){
	args := m.Called(id)
	return args.String(0), args.Error(1)
}
func (m *RsvUseCaseMock) GetApprovalList(page, size int) ([]dto.TransactionDTO, shared_model.Paging, error){
	args := m.Called(page, size)
	return args.Get(0).([]dto.TransactionDTO), args.Get(1).(shared_model.Paging), args.Error(2)
}
func (m *RsvUseCaseMock) GetAvailableRoom(payload dto.PayloadAvailable) ([]dto.RoomResponse, error){
	args := m.Called(payload)
	return args.Get(0).([]dto.RoomResponse), args.Error(1)
}