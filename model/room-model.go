package model

import "time"

type Room struct {
	Id         string     `json:"id"`
	CodeRoom   string     `json:"code_room"`
	RoomType   string     `json:"room_type"`
	Facilities string     `json:"facilities"`
	Capacity   int        `json:"capacity"`
	CreatedAt  time.Time  `json:"-"`
	UpdatedAt  time.Time  `json:"-"`
	DeletedAt  *time.Time `json:"-"`
}
