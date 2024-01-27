package usecase_mock

import (
	"booking-room/model"
	"booking-room/model/dto"
	"booking-room/shared/shared_model"

	"github.com/stretchr/testify/mock"
)

type FacilitiesUseCaseMock struct {
	mock.Mock
}

func (m *FacilitiesUseCaseMock) List(page, size int) ([]dto.FacilitiesResponse, shared_model.Paging, error) {
	args := m.Called(page, size)
	return args.Get(0).([]dto.FacilitiesResponse), args.Get(1).(shared_model.Paging), args.Error(2)
}
func (m *FacilitiesUseCaseMock) Get(id string) (model.Facilities, error) {
	args := m.Called(id)
	return args.Get(0).(model.Facilities), args.Error(1)
}
func (m *FacilitiesUseCaseMock) GetName(name string) (model.Facilities, error) {
	args := m.Called(name)
	return args.Get(0).(model.Facilities), args.Error(1)
}
func (m *FacilitiesUseCaseMock) GetType(ftype string, page, size int) ([]dto.FacilitiesResponse, shared_model.Paging, error) {
	args := m.Called(ftype, page, size)
	return args.Get(0).([]dto.FacilitiesResponse), args.Get(1).(shared_model.Paging), args.Error(2)
}
func (m *FacilitiesUseCaseMock) GetStatus(status string, page, size int) ([]dto.FacilitiesResponse, shared_model.Paging, error) {
	args := m.Called(status, page, size)
	return args.Get(0).([]dto.FacilitiesResponse), args.Get(1).(shared_model.Paging), args.Error(2)
}
func (m *FacilitiesUseCaseMock) Create(payload model.Facilities) (dto.FacilitiesCreated, error) {
	args := m.Called(payload)
	return args.Get(0).(dto.FacilitiesCreated), args.Error(1)
}
func (m *FacilitiesUseCaseMock) Update(payload model.Facilities, id string) (dto.FacilitiesUpdated, error) {
	args := m.Called(payload, id)
	return args.Get(0).(dto.FacilitiesUpdated), args.Error(1)
}
func (m *FacilitiesUseCaseMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
func (m *FacilitiesUseCaseMock) DeleteByName(name string) error {
	args := m.Called(name)
	return args.Error(0)
}

func (m *FacilitiesUseCaseMock) GetDeleted(page, size int) ([]model.Facilities, shared_model.Paging, error) {
	args := m.Called(page, size)
	return args.Get(0).([]model.Facilities), args.Get(1).(shared_model.Paging), args.Error(2)
}
