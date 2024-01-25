package model

import "time"

type EmployeeList struct {
	Name     string
	Email    string
	Division string
	Position string
	Role     string
	Contact  string
}
type RoomList struct {
	CodeRoom   string
	TypeRoom   string
	Capacity   int
	Facilities string
}

type FacilitiesList struct {
	CodeName     string
	FacilityType string
	Status       string
}

type ReservationReport struct {
	ReservationId       string
	EmployeeName        string
	CodeRoom            string
	StartDate           time.Time
	EndDate             time.Time
	Note                string
	ApproveStatus       string
	ApproveNote         string
	AdditionalFacilitys []AdditionalReport
}

type AdditionalReport struct {
	FacilitiesName string
}
type FacilityTotalReserved struct {
	FacilityType string
	Total        int
}

type RoomTotalReserved struct {
	RoomType string
	Total    int
}
