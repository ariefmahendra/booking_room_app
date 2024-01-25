package common

import (
	"booking-room/model"
	"booking-room/model/dto"
)

func EmployeeModelToResponse(employee model.EmployeeModel) dto.EmployeeResponse {
	return dto.EmployeeResponse{
		Id:        employee.Id,
		Name:      employee.Name,
		Email:     employee.Email,
		Division:  employee.Division,
		Position:  employee.Position,
		Role:      employee.Role,
		Contact:   employee.Contact,
		CreatedAt: employee.CreatedAt,
		UpdatedAt: employee.UpdatedAt,
		DeletedAt: employee.DeletedAt,
	}
}

func RequestToEmployeeModel(payload dto.EmployeeCreateRequest) model.EmployeeModel {
	return model.EmployeeModel{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
		Division: payload.Division,
		Position: payload.Position,
		Role:     payload.Role,
		Contact:  payload.Contact,
	}
}

func RoomModelToResponse(room model.Room) dto.RoomResponse {
	return dto.RoomResponse{
		Id:         room.Id,
		CodeRoom:   room.CodeRoom,
		RoomType:   room.RoomType,
		Facilities: room.Facilities,
		Capacity:   room.Capacity,
	}
}

func RequestToRoomModel(request dto.RoomRequest) model.Room {
	return model.Room{
		Id:         request.Id,
		CodeRoom:   request.CodeRoom,
		RoomType:   request.RoomType,
		Facilities: request.Facilities,
		Capacity:   request.Capacity,
	}
}
