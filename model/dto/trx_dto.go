package dto

import "time"

type TransactionDTO struct {
	Id         string `json:"id"`
	EmployeeId string `json:"employee_id"`
	EmplyName string `json:"employee_name"`
	RoomCode     string `json:"room_code"`
	StartDate  time.Time `json:"booked_start_date"`
	EndDate    time.Time `json:"booked_end_date"`
	Note string `json:"note"`
	ApproveStatus string `json:"approval_status"`
	ApproveNote string `json:"apprv_note"`
	Facility []Facility `json:"facilities"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	DeleteAt *time.Time `json:"delete_at"`
}

type Facility struct{
	Id string `json:"id"`
	Code string `json:"code"`
	Type string `json:"type"`
	
}