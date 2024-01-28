package repository

import (
	"booking-room/config"
	"booking-room/model"
	"booking-room/shared/shared_model"
	"database/sql"
	"math"
)

type RoomRepository interface {
	CreateRoom(payload model.Room) (model.Room, error)
	GetRoom(id string) (model.Room, error)
	UpdateRoom(payload model.Room) (model.Room, error)
	ListRoom(page, size int) ([]model.Room, shared_model.Paging, error)
}

type roomRepository struct {
	db *sql.DB
}

func (r *roomRepository) CreateRoom(payload model.Room) (model.Room, error) {
	var room model.Room

	err := r.db.QueryRow(config.CreateRoom, payload.CodeRoom, payload.RoomType, payload.Capacity, payload.Facilities).Scan(&room.Id, &room.CodeRoom, &room.RoomType, &room.Capacity, &room.Facilities, &room.CreatedAt, &room.UpdatedAt, &room.DeletedAt)

	if err != nil {
		return model.Room{}, err
	}

	return room, nil
}

func (r *roomRepository) UpdateRoom(payload model.Room) (model.Room, error) {
	var room model.Room

	err := r.db.QueryRow(config.UpdateRoomByID, room.Id, payload.CodeRoom, payload.RoomType, payload.Capacity, payload.Facilities).
		Scan(&room.Id, &room.CodeRoom, &room.RoomType, &room.Capacity, &room.Facilities, &room.CreatedAt, &room.UpdatedAt, &room.DeletedAt)

	if err != nil {
		return model.Room{}, err
	}

	return room, nil
}

func (r *roomRepository) GetRoom(id string) (model.Room, error) {
	var room model.Room
	err := r.db.QueryRow(config.SelectRoomByID, id).Scan(&room.Id, &room.CodeRoom, &room.RoomType, &room.Capacity, &room.Facilities, &room.CreatedAt, &room.UpdatedAt, &room.DeletedAt)

	if err != nil {
		return model.Room{}, err
	}

	return room, nil
}

func (r *roomRepository) ListRoom(page, size int) ([]model.Room, shared_model.Paging, error) {
	var rooms []model.Room
	offset := (page - 1) * size

	rows, err := r.db.Query(config.SelectRoomList, size, offset)
	if err != nil {
		return nil, shared_model.Paging{}, err
	}

	for rows.Next() {
		var room model.Room
		err := rows.Scan(&room.Id, &room.CodeRoom, &room.RoomType, &room.Capacity, &room.Facilities, &room.CreatedAt, &room.UpdatedAt, &room.DeletedAt)
		if err != nil {
			return nil, shared_model.Paging{}, err
		}

		rooms = append(rooms, room)
	}

	totalRows := 0
	if err := r.db.QueryRow(config.SelectCountRoom).Scan(&totalRows); err != nil {
		return nil, shared_model.Paging{}, err
	}

	paging := shared_model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}

	return rooms, paging, nil
}

func NewRoomRepository(db *sql.DB) RoomRepository {
	return &roomRepository{db: db}
}
