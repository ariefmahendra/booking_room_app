package usecase

import (
	"booking-room/model"
	"booking-room/repository"
	"booking-room/shared/shared_model"
	"fmt"
	//"strings"
	"time"
)

type RoomUseCase interface {
	RegisterNewRoom(payload model.Room) (model.Room, error)
	FindRoomById(id string) (model.Room, error)
	FindAllRoom(page, size int) ([]model.Room, shared_model.Paging, error)
	UpdateRoom(payload model.Room) (model.Room, error)
}

type roomUseCase struct {
	repo repository.RoomRepository
}

// FindRoomByID implements RoomUseCase.
func (r *roomUseCase) FindRoomById(id string) (model.Room, error) {
	return r.repo.GetRoom(id)
}

// FindAllRoom implements RoomUseCase.
func (r *roomUseCase) FindAllRoom(page, size int) ([]model.Room, shared_model.Paging, error) {
	return r.repo.ListRoom(page, size)
}


// RegisterNewRoom implements RoomUseCase.
func (r *roomUseCase) RegisterNewRoom(payload model.Room) (model.Room, error) {
    if payload.CodeRoom == "" || payload.RoomType == "" || payload.Capacity == 0 {
        return model.Room{}, fmt.Errorf("oops, field required")
    }

    if payload.Facilities == "" {
        payload.Facilities = "Standard Facilities"
    }

    room, err := r.repo.CreateRoom(payload)
    if err != nil {
        return model.Room{}, fmt.Errorf("failed to create a new room: %v", err)
    }

    return room, nil
}

// UpdateRoom implements RoomUseCase.
func (r *roomUseCase) UpdateRoom(payload model.Room) (model.Room, error) {
    if payload.Id == "" || payload.CodeRoom == "" || payload.RoomType == "" || payload.Capacity == 0 || payload.Facilities == ""{
        return model.Room{}, fmt.Errorf("oops, field required")
    }

    if payload.Facilities == "" {
        payload.Facilities = "Standard Facilities"
    }

    payload.UpdatedAt = time.Now()

    room, err := r.repo.UpdateRoom(payload)
    if err != nil {
        return model.Room{}, fmt.Errorf("failed to update room with ID %s: %v", payload.Id, err)
    }

    return room, nil
}

func NewRoomUseCase(repo repository.RoomRepository) RoomUseCase {
    return &roomUseCase{repo: repo}
}
