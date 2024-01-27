package model

import "time"

type Room struct {
	Id         string
	CodeRoom   string
	RoomType   string
	Facilities string
	Capacity   int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}
