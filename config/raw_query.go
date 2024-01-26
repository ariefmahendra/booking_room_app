package config

const (
	//room repository
	CreateRoom      = `INSERT INTO mst_room (code_room, room_type, capacity, facilities) VALUES ($1, $2, $3, $4) RETURNING id, code_room, room_type, facilities, capacity, created_at, updated_at, deleted_at`
	UpdateRoomByID  = `UPDATE mst_room SET code_room = $2, room_type = $3, capacity = $4, facilities = $5, updated_at = CURRENT_TIMESTAMP WHERE id = $1 RETURNING id, code_room, room_type, facilities, capacity, created_at, updated_at, deleted_at`
	SelectRoomByID  = `SELECT id, code_room, room_type, capacity, facilities, created_at, updated_at, deleted_at FROM mst_room WHERE id = $1`
	SelectRoomList  = `SELECT id, code_room, room_type, capacity, facilities, created_at, updated_at, deleted_at FROM mst_room ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	SelectCountRoom = `SELECT COUNT(*) FROM mst_room`

	//Employee Repository
	InsertEmployee        = `INSERT INTO mst_employee (name, email, password, division, position, role, contact) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, name, email, division, position, role, contact, created_at`
	GetEmployeeById       = `SELECT id, name, email, division, position, role, contact, created_at, updated_at, deleted_at FROM mst_employee WHERE id = $1 AND deleted_at IS NULL`
	GetEmployeeByEmail    = `SELECT id, name, email, division, position, role, contact, created_at, updated_at, deleted_at FROM mst_employee WHERE email = $1 AND deleted_at IS NULL`
	GetEmployees          = `SELECT id, name, email, division, position, role, contact, created_at, updated_at, deleted_at FROM mst_employee LIMIT $1 OFFSET $2 WHERE deleted_at IS NULL`
	GetDeletedEmployees   = `SELECT id, name, email, division, position, role, contact, created_at, updated_at, deleted_at FROM mst_employee WHERE deleted_at IS NOT NULL LIMIT $1 OFFSET $2`
	DeleteEmployeeById    = `UPDATE mst_employee SET deleted_at = (CURRENT_TIMESTAMP) WHERE id = $1`
	UpdateEmployeeById    = `UPDATE mst_employee SET name = $1, email = $2, password = $3, division = $4, position = $5, role = $6, contact = $7, updated_at = (CURRENT_TIMESTAMP) WHERE id = $8  AND  deleted_at IS NULL RETURNING id, name, email, division, position, role, contact, created_at, updated_at`
	PagingEmployeeActive  = `SELECT COUNT(*) FROM mst_employee WHERE deleted_at IS NULL`
	PagingEmployeeDeleted = `SELECT COUNT(*) FROM mst_employee`

	//Facility Repository
	RawPagingCount        = `"SELECT COUNT (*) FROM mst_facilities WHERE deleted_at IS NULL`
	FacilitiesList        = `SELECT code_name, facilities_type, status FROM mst_facilities WHERE deleted_at IS NULL LIMIT $1 OFFSET $2`
	FacilityGetId         = `SELECT * FROM mst_facilities WHERE id=$1 AND deleted_at IS NULL`
	FacilityGetName       = `SELECT * FROM mst_facilities WHERE code_name=$1 AND deleted_at IS NULL`
	FacilitiesCountStatus = `SELECT COUNT (*) FROM mst_facilities WHERE status=$1 AND deleted_at IS NULL`
	FacilitiesGetStatus   = `SELECT code_name, facilities_type, status FROM mst_facilities WHERE status=$1 AND deleted_at IS NULL LIMIT $2 OFFSET $3`
	FacilitiesCountType   = `SELECT COUNT (*) FROM mst_facilities WHERE facilities_type=$1 AND deleted_at IS NULL`
	FacilitiesGetType     = `SELECT code_name, facilities_type, status FROM mst_facilities WHERE facilities_type=$1 AND deleted_at IS NULL LIMIT $2 OFFSET $3`
	FacilityCountDeleted  = `SELECT COUNT (*) FROM mst_facilities WHERE deleted_at IS NOT NULL`
	FacilityGetDeleted    = `SELECT * FROM mst_facilities WHERE deleted_at IS NOT NULL LIMIT $1 OFFSET $2`
	FacilityInsert        = `INSERT INTO mst_facilities(code_name, facilities_type) VALUES($1, $2) RETURNING id, status, created_at`
	FacilityUpdate        = `UPDATE mst_facilities SET code_name=$1, facilities_type=$2, status=$3, updated_at=current_timestamp WHERE id=$4 RETURNING id, code_name, facilities_type, status, updated_at`
	FacilityDeleteById    = `UPDATE mst_facilities SET deleted_at=current_timestamp WHERE id=$1`
	FAcilityDeleteByName  = `UPDATE mst_facilities SET deleted_at=current_timestamp WHERE code_name=$1`
)
