package model

import "time"

type EmployeeModel struct {
	Id        string
	Email     string
	Name      string
	Password  string
	Division  string
	Position  string
	Role      string
	Contact   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
