package repository

import (
	"booking-room/model/dto"
	"booking-room/shared/shared_model"
	"database/sql"
	"log"
	"math"
)

type TrxRsvRepository interface {
	List(page,size int) ([]dto.TransactionDTO,  shared_model.Paging, error)
	GetID(id string) (dto.TransactionDTO, error)
	GetEmployee(id string, page,size int) ([]dto.TransactionDTO, shared_model.Paging, error)
}

type trxRsvRepository struct {
	db *sql.DB
}

// GetEmployee implements TrxRsvRepository.
func (t *trxRsvRepository) GetEmployee(id string, page,size int) ([]dto.TransactionDTO, shared_model.Paging, error) {
	var trxEmployee []dto.TransactionDTO
	offset := (page - 1) * size
	var err error

	query := `
			SELECT 
				tx.id, 
				tx.employee_id, 
				emp.name, 
				room.code_room, 
				tx.start_date, 
				tx.end_date,
				tx.notes,
				tx.approval_status,
				tx.approval_note,
				tx.created_at,
				tx.updated_at,
				tx.deleted_at
			FROM
				tx_room_reservation tx
			JOIN
				mst_employee emp ON tx.employee_id = emp.id
			JOIN
				mst_room room ON tx.room_id = room.id
			WHERE tx.employee_id = $1 AND tx.deleted_at IS NULL
			ORDER BY tx.created_at DESC
			LIMIT $2 OFFSET $3
			`

	row := t.db.QueryRow(query, id, size, offset)
	var emp dto.TransactionDTO
	if err := row.Scan(
			&emp.Id,
			&emp.EmployeeId,
			&emp.EmplyName,
			&emp.RoomCode,
			&emp.StartDate,
			&emp.EndDate,
			&emp.Note,
			&emp.ApproveStatus,
			&emp.ApproveNote,
			&emp.CreateAt,
			&emp.UpdateAt,
			&emp.DeleteAt,
	); err != nil {
		log.Println("trxRepository.QueryRow", err.Error())
		return []dto.TransactionDTO{}, shared_model.Paging{}, err
	}

	queryFacility := `
	SELECT a.id, f.code_name, f.facilities_type
	from tx_additional a
	JOIN mst_facilities f ON a.facilities_id = f.id
	WHERE a.reservation_id = $1
	`
	fclty, err := t.db.Query(queryFacility, emp.Id)
	if err != nil {
		return nil, shared_model.Paging{}, err
	}
	for fclty.Next() {
		var f dto.Facility
		err = fclty.Scan(&f.Id, &f.Code, &f.Type)
		if err != nil {
			return nil, shared_model.Paging{}, err
		}
		emp.Facility = append(emp.Facility, f)
	}
	trxEmployee = append(trxEmployee, emp)

	totalRows := 0
	queryRow := "SELECT COUNT(*) FROM tx_room_reservation WHERE employee_id = $1"
	if err := t.db.QueryRow(queryRow, id).Scan(&totalRows); err != nil {
		return nil, shared_model.Paging{}, err
	}

	paging := shared_model.Paging{
		Page:        page,
		RowsPerPage: page,
		TotalRows:   page,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}

	return trxEmployee, paging, err
}

// Get implements TrxRsvRepository.
func (t *trxRsvRepository) GetID(id string) (dto.TransactionDTO, error) {
	var trx dto.TransactionDTO
	var err error

	query := `SELECT 
				tx.id, 
				tx.employee_id, 
				emp.name, 
				room.code_room, 
				tx.start_date, 
				tx.end_date,
				tx.notes,
				tx.approval_status,
				tx.approval_note,
				tx.created_at,
				tx.updated_at,
				tx.deleted_at
			FROM
				tx_room_reservation tx
			JOIN
				mst_employee emp ON tx.employee_id = emp.id
			JOIN
				mst_room room ON tx.room_id = room.id
			WHERE tx.id = $1`

	row := t.db.QueryRow(query, id)
	if err := row.Scan(
		&trx.Id,
		&trx.EmployeeId,
		&trx.EmplyName,
		&trx.RoomCode,
		&trx.StartDate,
		&trx.EndDate,
		&trx.Note,
		&trx.ApproveStatus,
		&trx.ApproveNote,
		&trx.CreateAt,
		&trx.UpdateAt,
		&trx.DeleteAt,
	); err != nil {
		log.Println("trxRepository.QueryRow", err.Error())
		return dto.TransactionDTO{}, err
	}
	queryFacility := `
	SELECT a.id, f.code_name, f.facilities_type
	from tx_additional a
	JOIN mst_facilities f ON a.facilities_id = f.id
	WHERE a.reservation_id = $1
	`
	fclty, err := t.db.Query(queryFacility, trx.Id)
	if err != nil {
		return dto.TransactionDTO{}, err
	}
	for fclty.Next() {
		var f dto.Facility
		err = fclty.Scan(&f.Id, &f.Code, &f.Type)
		if err != nil {
			return dto.TransactionDTO{}, err
		}
		trx.Facility = append(trx.Facility, f)
	}

	return trx, err
}

// List implements TrxRsvRepository.
func (t *trxRsvRepository) List(page,size int) ([]dto.TransactionDTO, shared_model.Paging, error) {
	var trxList []dto.TransactionDTO
	offset := (page - 1) * size

	query := 
	`SELECT 
		tx.id, 
		tx.employee_id, 
		emp.name, 
		room.code_room, 
		tx.start_date, 
		tx.end_date,
		tx.notes,
		tx.approval_status,
		tx.approval_note,
		tx.created_at,
		tx.updated_at,
		tx.deleted_at
	FROM
		tx_room_reservation tx
	JOIN
		mst_employee emp ON tx.employee_id = emp.id
	JOIN
		mst_room room ON tx.room_id = room.id
	WHERE tx.deleted_at IS NULL
	ORDER BY tx.created_at DESC
	LIMIT $1 OFFSET $2;`

	rows, err := t.db.Query(query, size, offset)
	for rows.Next() {
		var trx dto.TransactionDTO
		err = rows.Scan(
			&trx.Id,
			&trx.EmployeeId,
			&trx.EmplyName,
			&trx.RoomCode,
			&trx.StartDate,
			&trx.EndDate,
			&trx.Note,
			&trx.ApproveStatus,
			&trx.ApproveNote,
			&trx.CreateAt,
			&trx.UpdateAt,
			&trx.DeleteAt,
		)
		if err != nil {
			log.Println("trxRepository.Scan", err.Error())
			return nil,shared_model.Paging{}, err
		}

		queryFacility := `
		SELECT a.id, f.code_name, f.facilities_type
		from tx_additional a
		JOIN mst_facilities f ON a.facilities_id = f.id
		WHERE a.reservation_id = $1
		`
		fclty, err := t.db.Query(queryFacility, trx.Id)
		if err != nil {
			return nil, shared_model.Paging{}, err
		}
		for fclty.Next() {
			var f dto.Facility
			err = fclty.Scan(&f.Id, &f.Code, &f.Type)
			if err != nil {
				return nil, shared_model.Paging{}, err
			}
			trx.Facility = append(trx.Facility, f)
		}

		trxList = append(trxList, trx)
	}

	totalRows := 0
	if err := t.db.QueryRow("SELECT COUNT(*) FROM tx_room_reservation").Scan(&totalRows); err != nil {
		return nil, shared_model.Paging{}, err
	}

	paging := shared_model.Paging{
		Page:        page,
		RowsPerPage: page,
		TotalRows:   page,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}

	return trxList, paging, err
}

func NewTrxRsvRepository(db *sql.DB) TrxRsvRepository {
	return &trxRsvRepository{
		db: db,
	}
}
