package dto

type RoomRequest struct {
	Id         string `json:"id"`
	CodeRoom   string `json:"code_room" json:"code_room,omitempty"`
	RoomType   string `json:"room_type" json:"room_type,omitempty"`
	Facilities string `json:"facilities" json:"facilities,omitempty"`
	Capacity   int    `json:"capacity" json:"capacity,omitempty"`
}

type RoomResponse struct {
	Id         string `json:"id"`
	CodeRoom   string `json:"code_room"`
	RoomType   string `json:"room_type"`
	Facilities string `json:"facilities"`
	Capacity   int    `json:"capacity"`
}

