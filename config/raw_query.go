package config

const (
<<<<<<< HEAD
	CreateRoom 			  = `INSERT INTO mst_room (code_room, room_type, capacity, facilities) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`
	UpdateRoomByID = `UPDATE mst_room SET code_room = $2, room_type = $3, capacity = $4, facilities = $5, updated_at = CURRENT_TIMESTAMP WHERE id = $1 RETURNING created_at, updated_at`
	SelectRoomByID        = `SELECT id, code_room, room_type, capacity, facilities, created_at, updated_at FROM mst_room WHERE id = $1`
	SelectRoomList        = `SELECT id, code_room, room_type, capacity, facilities, created_at, updated_at FROM mst_room ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	SelectCountRoom       = `SELECT COUNT(*) FROM mst_room`
)
=======
	RawQueryInsertEmployee     = `INSERT INTO mst_employee (name, email, password, division, position, role, contact) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, name, email, division, position, role, contact, created_at`
	RawQueryGetEmployeeById    = `SELECT id, name, email, division, position, role, contact, created_at, updated_at, deleted_at FROM mst_employee WHERE id = $1`
	RawQueryGetEmployeeByEmail = `SELECT id, name, email, division, position, role, contact, created_at, updated_at, deleted_at FROM mst_employee WHERE email = $1`
	RawQueryGetEmployees       = `SELECT id, name, email, division, position, role, contact, created_at, updated_at, deleted_at FROM mst_employee LIMIT $1 OFFSET $2`
	RawDeleteEmployeeById      = `UPDATE mst_employee SET deleted_at = (CURRENT_TIMESTAMP) WHERE id = $1`
	RawUpdateEmployeeById      = `UPDATE mst_employee SET name = $1, email = $2, password = $3, division = $4, position = $5, role = $6, contact = $7, updated_at = (CURRENT_TIMESTAMP) WHERE id = $8 RETURNING id, name, email, division, position, role, contact, created_at, updated_at`
	RAWQueryPaging             = `SELECT COUNT (*) FROM mst_employee`
)
>>>>>>> ed7e6ada7c231957f8498b03fd926752a5f88f1d
