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
