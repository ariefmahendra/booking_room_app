package dto

import "time"

type EmployeeResponse struct {
	Id        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email,omitempty"`
	Division  string     `json:"division"`
	Position  string     `json:"position"`
	Role      string     `json:"role"`
	Contact   string     `json:"contact"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
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
