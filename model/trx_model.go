package model

import "time"

type Transaction struct {
	Id         string `json:"id"`
	EmployeeId string `json:"employee_id"`
	RoomId     string `json:"room_id"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	Note string `json:"note"`
	ApproveStatus string `json:"approval_status"`
	ApproveNote string `json:"apprv_note"`
	Facility []Facility `json:"facilities"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type Facility struct{
	Id string `json:"id"`
	Code string `json:"code"`
	Type string `json:"type"`
}