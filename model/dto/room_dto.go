package dto

type RoomRequest struct {
	Id         string `json:"id"`
	CodeRoom   string `json:"code_room"`
	RoomType   string `json:"room_type"`
	Facilities string `json:"facilities"`
	Capacity   int    `json:"capacity"`
}

type RoomResponse struct {
	Id         string `json:"id"`
	CodeRoom   string `json:"code_room"`
	RoomType   string `json:"room_type"`
	Facilities string `json:"facilities"`
	Capacity   int    `json:"capacity"`
}
