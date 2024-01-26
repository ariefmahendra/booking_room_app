package dto

type FacilitiesCreated struct {
	Id             string `json:"id"`
	CodeName       string `json:"codeName"`
	FacilitiesType string `json:"FacilitiesType"`
	Status         string `json:"status"`
	CreatedAt      string `json:"createdAt"`
}
type FacilitiesUpdated struct {
	Id             string `json:"id"`
	CodeName       string `json:"codeName"`
	FacilitiesType string `json:"FacilitiesType"`
	Status         string `json:"status"`
	UpdatedAt      string `json:"updatedAt"`
}

type FacilitiesResponse struct {
	CodeName       string `json:"codeName"`
	FacilitiesType string `json:"FacilitiesType"`
	Status         string `json:"status"`
}
