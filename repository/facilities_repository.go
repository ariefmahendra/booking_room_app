package repository

import (
	"booking-room/config"
	"booking-room/model"
	"booking-room/model/dto"
	"booking-room/shared/shared_model"
	"database/sql"
	"log"
	"math"

	_ "github.com/lib/pq"
)

type FacilitiesRepository interface {
	List(page, size int) ([]dto.FacilitiesResponse, shared_model.Paging, error)
	Get(id string) (model.Facilities, error)
	GetName(name string) (model.Facilities, error)
	GetStatus(status string, page, size int) ([]dto.FacilitiesResponse, shared_model.Paging, error)
	GetType(status string, page, size int) ([]dto.FacilitiesResponse, shared_model.Paging, error)
	Create(payload model.Facilities) (dto.FacilitiesCreated, error)
	Update(facility model.Facilities, id string) (dto.FacilitiesUpdated, error)
	Delete(id string) error
	DeleteByName(name string) error
	GetDeleted(page, size int) ([]model.Facilities, shared_model.Paging, error)
}

var deletedAt sql.NullString

type facilitiesRepository struct {
	db *sql.DB
}

// Query all facilites paged
func (f *facilitiesRepository) List(page, size int) ([]dto.FacilitiesResponse, shared_model.Paging, error) {
	//set max page
	totalRows := 0
	if err := f.db.QueryRow(config.RawPagingCount).Scan(&totalRows); err != nil {
		log.Println("facilitiesRepository.QueryRow count", err.Error())
		return nil, shared_model.Paging{}, err
	}
	paging := shared_model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}
	if page > paging.TotalPages {
		page = paging.TotalPages
		paging.Page = page
	}
	offset := (page - 1) * size

	rows, err := f.db.Query(config.FacilitiesList, size, offset)
	if err != nil {
		log.Println("facilitiesRepository.Query", err.Error())
		return nil, shared_model.Paging{}, err
	}
	defer rows.Close()

	// append all facilities data into facilities struct
	var facilities []dto.FacilitiesResponse
	for rows.Next() {
		var facility dto.FacilitiesResponse
		if err := rows.Scan(&facility.CodeName, &facility.FacilitiesType, &facility.Status); err != nil {
			log.Println("facilitiesRepository.Scan", err.Error())
			return nil, shared_model.Paging{}, err
		}

		facilities = append(facilities, facility)
		if err = rows.Err(); err != nil {
			return nil, shared_model.Paging{}, err
		}
	}
	return facilities, paging, nil
}

// Query facility by id
func (f *facilitiesRepository) Get(id string) (model.Facilities, error) {
	var facility model.Facilities
	if err := f.db.QueryRow(config.FacilityGetId, id).Scan(&facility.Id, &facility.CodeName, &facility.FacilitiesType, &facility.Status, &facility.CreatedAt, &facility.UpdatedAt, &deletedAt); err != nil {
		log.Println("facilitiesRepository.Get", err.Error())
		return model.Facilities{}, err
	}
	return facility, nil
}

// Query facility by name
func (f *facilitiesRepository) GetName(name string) (model.Facilities, error) {
	var facility model.Facilities
	if err := f.db.QueryRow(config.FacilityGetName, name).Scan(&facility.Id, &facility.CodeName, &facility.FacilitiesType, &facility.Status, &facility.CreatedAt, &facility.UpdatedAt, &deletedAt); err != nil {
		log.Println("facilitiesRepository.Get", err.Error())
		return model.Facilities{}, err
	}
	return facility, nil
}

// Query facility by status
func (f *facilitiesRepository) GetStatus(status string, page, size int) ([]dto.FacilitiesResponse, shared_model.Paging, error) {
	var facilities []dto.FacilitiesResponse
	//set max page
	totalRows := 0
	if err := f.db.QueryRow(config.FacilitiesCountStatus, status).Scan(&totalRows); err != nil {
		return nil, shared_model.Paging{}, err
	}
	paging := shared_model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}
	if page > paging.TotalPages {
		page = paging.TotalPages
		paging.Page = page
	}
	offset := (page - 1) * size

	rows, err := f.db.Query(config.FacilitiesGetStatus, status, size, offset)
	if err != nil {
		log.Println("facilitiesRepository.Query", err.Error())
		return nil, shared_model.Paging{}, err
	}
	defer rows.Close()

	// append all facilities data into facilities struct
	for rows.Next() {
		var facility dto.FacilitiesResponse
		if err := rows.Scan(&facility.CodeName, &facility.FacilitiesType, &facility.Status); err != nil {
			log.Println("facilitiesRepository.Scan", err.Error())
			return nil, shared_model.Paging{}, err
		}

		facilities = append(facilities, facility)
		if err = rows.Err(); err != nil {
			return nil, shared_model.Paging{}, err
		}
	}
	return facilities, paging, nil
}

// Query facility by type
func (f *facilitiesRepository) GetType(ftype string, page, size int) ([]dto.FacilitiesResponse, shared_model.Paging, error) {
	var facilities []dto.FacilitiesResponse
	//set max page
	totalRows := 0
	if err := f.db.QueryRow(config.FacilitiesCountType, ftype).Scan(&totalRows); err != nil {
		return nil, shared_model.Paging{}, err
	}
	paging := shared_model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}
	if page > paging.TotalPages {
		page = paging.TotalPages
		paging.Page = page
	}
	offset := (page - 1) * size

	rows, err := f.db.Query(config.FacilitiesGetType, ftype, size, offset)
	if err != nil {
		log.Println("facilitiesRepository.Query", err.Error())
		return nil, shared_model.Paging{}, err
	}
	defer rows.Close()

	// append all facilities data into facilities struct
	for rows.Next() {
		var facility dto.FacilitiesResponse
		if err := rows.Scan(&facility.CodeName, &facility.FacilitiesType, &facility.Status); err != nil {
			log.Println("facilitiesRepository.Scan", err.Error())
			return nil, shared_model.Paging{}, err
		}

		facilities = append(facilities, facility)
		if err = rows.Err(); err != nil {
			return nil, shared_model.Paging{}, err
		}
	}
	return facilities, paging, nil
}

// Get deleted facility
func (f *facilitiesRepository) GetDeleted(page, size int) ([]model.Facilities, shared_model.Paging, error) {
	//set max page
	totalRows := 0
	if err := f.db.QueryRow(config.FacilityCountDeleted).Scan(&totalRows); err != nil {
		return nil, shared_model.Paging{}, err
	}
	paging := shared_model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}
	if page > paging.TotalPages {
		page = paging.TotalPages
		paging.Page = page
	}
	offset := (page - 1) * size

	rows, err := f.db.Query(config.FacilityGetDeleted, size, offset)
	if err != nil {
		log.Println("facilitiesRepository.Query", err.Error())
		return nil, shared_model.Paging{}, err
	}
	defer rows.Close()

	// append all facilities data into facilities struct
	var facilities []model.Facilities
	for rows.Next() {
		var facility model.Facilities
		if err := rows.Scan(&facility.Id, &facility.CodeName, &facility.FacilitiesType, &facility.Status, &facility.CreatedAt, &facility.UpdatedAt, &deletedAt); err != nil {
			log.Println("facilitiesRepository.Scan", err.Error())
			return nil, shared_model.Paging{}, err
		}

		facilities = append(facilities, facility)
		if err = rows.Err(); err != nil {
			return nil, shared_model.Paging{}, err
		}
	}
	return facilities, paging, nil
}

// Query to create new facility
func (f *facilitiesRepository) Create(payload model.Facilities) (dto.FacilitiesCreated, error) {
	var facility dto.FacilitiesCreated
	err := f.db.QueryRow(config.FacilityInsert, payload.CodeName, payload.FacilitiesType).Scan(&facility.Id, &facility.Status, &facility.CreatedAt)
	if err != nil {
		log.Println("facilitiesRepository.Create", err.Error())
		return dto.FacilitiesCreated{}, err
	}
	facility.CodeName = payload.CodeName
	facility.FacilitiesType = payload.FacilitiesType
	return facility, nil
}

// Query to update facility
func (f *facilitiesRepository) Update(payload model.Facilities, id string) (dto.FacilitiesUpdated, error) {
	var facility dto.FacilitiesUpdated
	err := f.db.QueryRow(config.FacilityUpdate,
		payload.CodeName, payload.FacilitiesType, payload.Status, id).Scan(&facility.Id, &facility.CodeName, &facility.FacilitiesType, &facility.Status, &facility.UpdatedAt)
	if err != nil {
		log.Println("facilitiesRepository.Update", err.Error())
		return dto.FacilitiesUpdated{}, err
	}

	return facility, nil
}

// Query to delete facility by id
func (f *facilitiesRepository) Delete(id string) error {
	_, err := f.db.Exec(config.FacilityDeleteById, id)
	if err != nil {
		log.Println("facilitiesRepository.Delete", err.Error())
	}
	return nil
}

// Query to delete facility by name
func (f *facilitiesRepository) DeleteByName(name string) error {
	_, err := f.db.Exec(config.FAcilityDeleteByName, name)
	if err != nil {
		log.Println("facilitiesRepository.Delete", err.Error())
	}
	return nil
}

// constructor for facilities repository
func NewFacilitiesRepository(db *sql.DB) FacilitiesRepository {
	return &facilitiesRepository{
		db: db,
	}
}
