package repo_mock

import (
	"booking-room/model/dto"
	"booking-room/shared/shared_model"

	"github.com/stretchr/testify/mock"
)

type RsvRepoMock struct {
	mock.Mock
}

func (m *RsvRepoMock) List(page, size int) ([]dto.TransactionDTO, shared_model.Paging, error){
	args := m.Called(page, size)
	return args.Get(0).([]dto.TransactionDTO), args.Get(1).(shared_model.Paging), args.Error(2)
}
func (m *RsvRepoMock) GetID(id string) (dto.TransactionDTO, error){
	args := m.Called(id)
	return args.Get(0).(dto.TransactionDTO), args.Error(1)
}
func (m *RsvRepoMock) GetEmployee(id string, page, size int) ([]dto.TransactionDTO, shared_model.Paging, error){
	args := m.Called(id, page, size)
	return args.Get(0).([]dto.TransactionDTO), args.Get(1).(shared_model.Paging), args.Error(2)
}
func (m *RsvRepoMock) PostReservation(payload dto.PayloadReservationDTO) (string, error){
	args := m.Called(payload)
	return args.String(0), args.Error(1)
}
func (m *RsvRepoMock) UpdateStatus(payload dto.TransactionDTO) (dto.TransactionDTO, error){
	args := m.Called(payload)
	return args.Get(0).(dto.TransactionDTO), args.Error(1)
}
func (m *RsvRepoMock) DeleteResv(id string) (string, error){
	args := m.Called(id)
	return args.String(0), args.Error(1)
}