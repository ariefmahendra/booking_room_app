package model

import "time"

type Room struct {
	Id         string     `json:"id"`
	CodeRoom   string     `json:"code_room"`
	RoomType   string     `json:"room_type"`
	Facilities string     `json:"facilities"`
	Capacity   int        `json:"capacity"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}
