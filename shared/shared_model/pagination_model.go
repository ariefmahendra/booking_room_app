package shared_model

type Paging struct {
<<<<<<< HEAD
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
=======
	Page       int `json:"page"`
	RowPerPage int `json:"row_per_page"`
	TotalRows  int `json:"total_rows"`
	TotalPages int `json:"total_pages"`
}
>>>>>>> ed7e6ada7c231957f8498b03fd926752a5f88f1d
