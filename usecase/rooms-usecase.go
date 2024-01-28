package usecase

import (
	"booking-room/model/dto"
	"booking-room/repository"
	"booking-room/shared/common"
	"booking-room/shared/shared_model"
	"fmt"
)

type RoomUseCase interface {
	RegisterNewRoom(payload dto.RoomRequest) (dto.RoomResponse, error)
	FindRoomById(id string) (dto.RoomResponse, error)
	FindAllRoom(page, size int) ([]dto.RoomResponse, shared_model.Paging, error)
	UpdateRoom(payload dto.RoomRequest) (dto.RoomResponse, error)
}

type roomUseCase struct {
	repo repository.RoomRepository
}

func (r *roomUseCase) FindRoomById(id string) (dto.RoomResponse, error) {
	room, err := r.repo.GetRoom(id)
	if err != nil {
		return dto.RoomResponse{}, err
	}

	return common.RoomModelToResponse(room), nil
}

func (r *roomUseCase) FindAllRoom(page, size int) ([]dto.RoomResponse, shared_model.Paging, error) {
	rooms, paging, err := r.repo.ListRoom(page, size)
	if err != nil {
		return nil, shared_model.Paging{}, err
	}

	var roomsResponse []dto.RoomResponse
	for _, room := range rooms {
		var roomResponse dto.RoomResponse
		roomResponse = common.RoomModelToResponse(room)
		roomsResponse = append(roomsResponse, roomResponse)
	}

	return roomsResponse, paging, nil
}

func (r *roomUseCase) RegisterNewRoom(payload dto.RoomRequest) (dto.RoomResponse, error) {
	room := common.RequestToRoomModel(payload)

	if room.CodeRoom == "" || room.RoomType == "" || room.Capacity == 0 {
		return dto.RoomResponse{}, fmt.Errorf("oops, field required")
	}

	if payload.Facilities == "" {
		payload.Facilities = "Standard Facilities"
	}

	room, err := r.repo.CreateRoom(room)
	if err != nil {
		return dto.RoomResponse{}, err
	}

	return common.RoomModelToResponse(room), nil
}

func (r *roomUseCase) UpdateRoom(payload dto.RoomRequest) (dto.RoomResponse, error) {
	room := common.RequestToRoomModel(payload)

	if room.Id == "" || room.CodeRoom == "" || room.RoomType == "" || room.Capacity == 0 || room.Facilities == "" {
		return dto.RoomResponse{}, fmt.Errorf("oops, field required")
	}

	if payload.Facilities == "" {
		payload.Facilities = "Standard Facilities"
	}

	room, err := r.repo.UpdateRoom(room)
	if err != nil {
		return dto.RoomResponse{}, err
	}

	return common.RoomModelToResponse(room), nil
}

func NewRoomUseCase(repo repository.RoomRepository) RoomUseCase {
	return &roomUseCase{repo: repo}
}
