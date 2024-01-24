package usecase

import (
	"booking-room/model"
	"booking-room/repository"
	"booking-room/shared/shared_model"
	"fmt"
	"strings"
)

type FacilitiesUsecase interface {
	List() ([]model.Facilities, error)
	ListPaged(page, size int) ([]model.Facilities, shared_model.Paging, error)
	Get(id string) (model.Facilities, error)
	GetByName(name string) (model.Facilities, error)
	GetByType(ftype string, page, size int) ([]model.Facilities, shared_model.Paging, error)
	GetByStatus(status string, page, size int) ([]model.Facilities, shared_model.Paging, error)
	Create(payload model.Facilities) (model.Facilities, error)
	Update(payload model.Facilities, id string) (model.Facilities, error)
	Delete(id string) error
	DeleteByName(name string) error
}

type facilitiesUsecase struct {
	facilitiesRepository repository.FacilitiesRepository
}

// usecase for geting all facilities
func (f *facilitiesUsecase) List() ([]model.Facilities, error) {
	facilites, err := f.facilitiesRepository.List()
	if err != nil {
		return nil, fmt.Errorf("Problem with accesing Facilities Data")
	}
	return facilites, err
}

// usecase for geting all facilities paged
func (f *facilitiesUsecase) ListPaged(page, size int) ([]model.Facilities, shared_model.Paging, error) {
	facilites, paging, err := f.facilitiesRepository.ListPaged(page, size)
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
func (f *facilitiesUsecase) GetByStatus(status string, page, size int) ([]model.Facilities, shared_model.Paging, error) {
	facility, paging, err := f.facilitiesRepository.GetStatus(status, page, size)
	if err != nil {
		return []model.Facilities{}, shared_model.Paging{}, fmt.Errorf("Status not found")
	}
	return facility, paging, err
}

// usecase for selecting facility by type
func (f *facilitiesUsecase) GetByType(ftype string, page, size int) ([]model.Facilities, shared_model.Paging, error) {
	facility, paging, err := f.facilitiesRepository.GetType(ftype, page, size)
	if err != nil {
		return []model.Facilities{}, shared_model.Paging{}, fmt.Errorf("Facility type not found")
	}
	return facility, paging, err
}

// usecase for creating new facility
func (f *facilitiesUsecase) Create(payload model.Facilities) (model.Facilities, error) {
	if payload.CodeName == "" || payload.FacilitiesType == "" {
		return model.Facilities{}, fmt.Errorf("Facility Code Name and Facility Facilities Type cannot be empty")
	}
	payload.CodeName = strings.ToUpper(payload.CodeName)
	payload.FacilitiesType = strings.ToLower(payload.FacilitiesType)
	facility, err := f.facilitiesRepository.Create(payload)
	if err != nil {
		return model.Facilities{}, fmt.Errorf("Failed to register new facility")
	}
	return facility, nil
}

// usecase for updating facility
func (f *facilitiesUsecase) Update(payload model.Facilities, id string) (model.Facilities, error) {
	payload.CodeName = strings.ToUpper(payload.CodeName)
	payload.FacilitiesType = strings.ToLower(payload.FacilitiesType)

	//check if id exist
	oldFacility, err := f.facilitiesRepository.Get(id)
	if err != nil {
		return model.Facilities{}, fmt.Errorf("Id not found")
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
		return model.Facilities{}, fmt.Errorf("You need to insert at least one input")
	}
	// check if input the same as old value
	if payload.CodeName == oldFacility.CodeName && payload.FacilitiesType == oldFacility.FacilitiesType && payload.Status == oldFacility.Status {
		return model.Facilities{}, fmt.Errorf("No changes detected")
	}
	//updating facility
	facility, err := f.facilitiesRepository.Update(payload, id)
	if err != nil {
		/* if err.Code == "23505" {
			return model.Facilities{}, fmt.Errorf("Facility Code Name or Facilities Type already exist")
		} */
		return model.Facilities{}, fmt.Errorf("Failed to update facility")
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
	return fmt.Errorf("Facility deleted")
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
	return fmt.Errorf("Facility deleted")
}

// constructor for facilities usecase
func NewFacilitiesUsecase(facilitiesRepository repository.FacilitiesRepository) FacilitiesUsecase {
	return &facilitiesUsecase{
		facilitiesRepository: facilitiesRepository,
	}
}
