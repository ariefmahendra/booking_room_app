package config

const (
	CreateRoom 			  = `INSERT INTO mst_room (code_room, room_type, capacity, facilities) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`
	UpdateRoomByID = `UPDATE mst_room SET code_room = $2, room_type = $3, capacity = $4, facilities = $5, updated_at = CURRENT_TIMESTAMP WHERE id = $1 RETURNING created_at, updated_at`
	SelectRoomByID        = `SELECT id, code_room, room_type, capacity, facilities, created_at, updated_at FROM mst_room WHERE id = $1`
	SelectRoomList        = `SELECT id, code_room, room_type, capacity, facilities, created_at, updated_at FROM mst_room ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	SelectCountRoom       = `SELECT COUNT(*) FROM mst_room`
)