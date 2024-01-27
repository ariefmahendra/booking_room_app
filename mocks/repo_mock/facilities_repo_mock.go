package repo_mock

import (
	"booking-room/model"
	"booking-room/model/dto"
	"booking-room/shared/shared_model"

	"github.com/stretchr/testify/mock"
)

type FacilitiesMock struct {
	mock.Mock
}

func (f *FacilitiesMock) List(page, size int) ([]dto.FacilitiesResponse, shared_model.Paging, error) {
	args := f.Called(page, size)
	return args.Get(0).([]dto.FacilitiesResponse), args.Get(1).(shared_model.Paging), args.Error(2)
}

func (f *FacilitiesMock) GetDeleted(page, size int) ([]model.Facilities, shared_model.Paging, error) {
	args := f.Called(page, size)
	return args.Get(0).([]model.Facilities), args.Get(1).(shared_model.Paging), args.Error(2)
}

func (f *FacilitiesMock) Get(id string) (model.Facilities, error) {
	args := f.Called(id)
	return args.Get(0).(model.Facilities), args.Error(1)
}

func (f *FacilitiesMock) GetName(name string) (model.Facilities, error) {
	args := f.Called(name)
	return args.Get(0).(model.Facilities), args.Error(1)
}

func (f *FacilitiesMock) GetStatus(status string, page, size int) ([]dto.FacilitiesResponse, shared_model.Paging, error) {
	args := f.Called(status, page, size)
	return args.Get(0).([]dto.FacilitiesResponse), args.Get(1).(shared_model.Paging), args.Error(2)
}

func (f *FacilitiesMock) GetType(ftype string, page, size int) ([]dto.FacilitiesResponse, shared_model.Paging, error) {
	args := f.Called(ftype, page, size)
	return args.Get(0).([]dto.FacilitiesResponse), args.Get(1).(shared_model.Paging), args.Error(2)
}

func (f *FacilitiesMock) Create(payload model.Facilities) (dto.FacilitiesCreated, error) {
	args := f.Called(payload)
	return args.Get(0).(dto.FacilitiesCreated), args.Error(1)
}

func (f *FacilitiesMock) Update(payload model.Facilities, id string) (dto.FacilitiesUpdated, error) {
	args := f.Called(payload, id)
	return args.Get(0).(dto.FacilitiesUpdated), args.Error(1)
}

func (f *FacilitiesMock) Delete(id string) error {
	args := f.Called(id)
	return args.Error(0)
}

func (f *FacilitiesMock) DeleteByName(name string) error {
	args := f.Called(name)
	return args.Error(0)
}
