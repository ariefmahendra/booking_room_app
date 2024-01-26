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
	CreateAt time.Time `json:"-"`
	UpdateAt time.Time `json:"-"`
	DeleteAt *time.Time `json:"-"`
}

type Facility struct{
	Id string `json:"id"`
	Code string `json:"code"`
	Type string `json:"type"`	
}

type PayloadReservationDTO struct {
	Id         string `json:"id"`
	Email 		string `json:"email"`
	RoomCode    string `json:"room_code"`
	StartDate  	*time.Time `json:"booked_start_date"`
	EndDate    	*time.Time `json:"booked_end_date"`
	Note 		string `json:"note"`
	Facilities 	[]Facility `json:"facilities"`
}

type PayloadAvailable struct{
	IdRoom string `json:"id_room"`
	StartDate string `json:"start_date"`
	EndDate string `json:"end_date"`
}