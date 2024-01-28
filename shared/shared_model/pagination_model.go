package shared_model

type Paging struct {
	Page        int `json:"page"`
	RowsPerPage int `json:"rowsPerPage"`
	TotalPages  int `json:"totalPages"`
	TotalRows   int `json:"totalRows"`
}

type ListResponse struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data"`
	Paging Paging      `json:"paging"`
}