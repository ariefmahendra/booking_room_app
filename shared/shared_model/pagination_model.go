package shared_model

type Paging struct {
	Page        int `json:"page"`
	TotalPages  int `json:"totalPages"`
	TotalRows   int `json:"totalRows"`
	RowsPerPage int `json:"rowsPerPage"`
}

type ListResponse struct {
	Status Status `json:"status"`
	Data   interface{} `json:"data"`
	Paging Paging `json:"paging"`
}
