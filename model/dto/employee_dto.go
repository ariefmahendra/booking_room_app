package dto

import "time"

type EmployeeResponse struct {
	Id        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Division  string     `json:"division"`
	Position  string     `json:"position"`
	Role      string     `json:"role"`
	Contact   string     `json:"contact"`
	CreatedAt time.Time  `json:"created_At"`
	UpdatedAt time.Time  `json:"updated_At"`
	DeletedAt *time.Time `json:"deleted_At"`
}

type EmployeeCreateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Division string `json:"division"`
	Position string `json:"position"`
	Role     string `json:"role"`
	Contact  string `json:"contact"`
}
