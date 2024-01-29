package usecase

import (
	"booking-room/model"
	"booking-room/model/dto"
	"booking-room/repository"
	"booking-room/shared/shared_model"
	"fmt"
	"strings"
)

type FacilitiesUsecase interface {
	List(page, size int) ([]dto.FacilitiesResponse, shared_model.Paging, error)
	Get(id string) (model.Facilities, error)
	GetByName(name string) (model.Facilities, error)
	GetByType(ftype string, page, size int) ([]dto.FacilitiesResponse, shared_model.Paging, error)
	GetByStatus(status string, page, size int) ([]dto.FacilitiesResponse, shared_model.Paging, error)
	Create(payload model.Facilities) (dto.FacilitiesCreated, error)
	Update(payload model.Facilities, id string) (dto.FacilitiesUpdated, error)
	Delete(id string) error
	DeleteByName(name string) error
	GetDeleted(page, size int) ([]model.Facilities, shared_model.Paging, error)
}

type facilitiesUsecase struct {
	facilitiesRepository repository.FacilitiesRepository
}

// usecase for geting all facilities paged
func (f *facilitiesUsecase) List(page, size int) ([]dto.FacilitiesResponse, shared_model.Paging, error) {
	facilites, paging, err := f.facilitiesRepository.List(page, size)
	if err != nil {
		return nil, shared_model.Paging{}, fmt.Errorf("Problem with accesing Facilities Data")
	}
	return facilites, paging, err
}

// usecase for selecting facility by id
func (f *facilitiesUsecase) Get(id string) (model.Facilities, error) {
	facility, err := f.facilitiesRepository.Get(id)
	if err != nil {
		return model.Facilities{}, fmt.Errorf("Id not found")
	}
	return facility, err
}

// usecase for selecting facility by name
func (f *facilitiesUsecase) GetByName(name string) (model.Facilities, error) {
	facility, err := f.facilitiesRepository.GetName(name)
	if err != nil {
		return model.Facilities{}, fmt.Errorf("Name not found")
	}
	return facility, err
}

// usecase for selecting facility by status
func (f *facilitiesUsecase) GetByStatus(status string, page, size int) ([]dto.FacilitiesResponse, shared_model.Paging, error) {
	facility, paging, err := f.facilitiesRepository.GetStatus(status, page, size)
	if err != nil {
		return []dto.FacilitiesResponse{}, shared_model.Paging{}, fmt.Errorf("Status not found")
	}
	return facility, paging, err
}

// usecase for selecting facility by type
func (f *facilitiesUsecase) GetByType(ftype string, page, size int) ([]dto.FacilitiesResponse, shared_model.Paging, error) {
	facility, paging, err := f.facilitiesRepository.GetType(ftype, page, size)
	if err != nil {
		return []dto.FacilitiesResponse{}, shared_model.Paging{}, fmt.Errorf("Facility type not found")
	}
	return facility, paging, err
}

// usecase for creating new facility
func (f *facilitiesUsecase) Create(payload model.Facilities) (dto.FacilitiesCreated, error) {
	payload.CodeName = strings.ToUpper(payload.CodeName)
	payload.FacilitiesType = strings.ToLower(payload.FacilitiesType)
	facility, err := f.facilitiesRepository.Create(payload)
	if err != nil {
		return dto.FacilitiesCreated{}, fmt.Errorf("Code name %s already exist", payload.CodeName)
	}
	return facility, nil
}

// usecase for updating facility
func (f *facilitiesUsecase) Update(payload model.Facilities, id string) (dto.FacilitiesUpdated, error) {
	payload.CodeName = strings.ToUpper(payload.CodeName)
	payload.FacilitiesType = strings.ToLower(payload.FacilitiesType)

	//check if id exist
	oldFacility, err := f.facilitiesRepository.Get(id)
	if err != nil {
		return dto.FacilitiesUpdated{}, fmt.Errorf("Id not found")
	}
	//chek if input empty, then old value will be used
	if payload.CodeName == "" {
		payload.CodeName = oldFacility.CodeName
	}
	if payload.FacilitiesType == "" {
		payload.FacilitiesType = oldFacility.FacilitiesType
	}
	if payload.Status == "" {
		payload.Status = oldFacility.Status
	}
	//check if input not empty
	if payload.CodeName == "" && payload.FacilitiesType == "" && payload.Status == "" {
		return dto.FacilitiesUpdated{}, fmt.Errorf("You need to insert at least one input")
	}
	// check if input the same as old value
	if payload.CodeName == oldFacility.CodeName && payload.FacilitiesType == oldFacility.FacilitiesType && payload.Status == oldFacility.Status {
		return dto.FacilitiesUpdated{}, fmt.Errorf("No changes detected")
	}
	//updating facility
	facility, err := f.facilitiesRepository.Update(payload, id)
	if err != nil {
		/* if err.Code == "23505" {
			return model.Facilities{}, fmt.Errorf("Facility Code Name or Facilities Type already exist")
		} */
		return dto.FacilitiesUpdated{}, fmt.Errorf("Failed to update facility, Code Name need to be unique")
	}
	return facility, nil
}

// delete facility by id
func (f *facilitiesUsecase) Delete(id string) error {
	_, err := f.facilitiesRepository.Get(id)
	if err != nil {
		return fmt.Errorf("Id not found")
	}
	if err := f.facilitiesRepository.Delete(id); err != nil {
		return fmt.Errorf("Failed to delete facility")
	}
	return nil
}

// delete facility by name
func (f *facilitiesUsecase) DeleteByName(name string) error {
	_, err := f.facilitiesRepository.GetName(name)
	if err != nil {
		return fmt.Errorf("Name not found")
	}
	if err := f.facilitiesRepository.DeleteByName(name); err != nil {
		return fmt.Errorf("Failed to delete facility")
	}
	return nil
}

// Get deleted facilities
func (f *facilitiesUsecase) GetDeleted(page, size int) ([]model.Facilities, shared_model.Paging, error) {
	facility, paging, err := f.facilitiesRepository.GetDeleted(page, size)
	if err != nil {
		return []model.Facilities{}, shared_model.Paging{}, fmt.Errorf("Facility not found")
	}
	return facility, paging, err
}

// constructor for facilities usecase
func NewFacilitiesUsecase(facilitiesRepository repository.FacilitiesRepository) FacilitiesUsecase {
	return &facilitiesUsecase{
		facilitiesRepository: facilitiesRepository,
	}
}
