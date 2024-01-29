package model

type Facilities struct {
	Id             string `json:"id"`
	CodeName       string `json:"codeName"`
	FacilitiesType string `json:"facilitiesType"`
	Status         string `json:"status"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
	DeletedAt      string `json:"deletedAt,omitempty"`
}
