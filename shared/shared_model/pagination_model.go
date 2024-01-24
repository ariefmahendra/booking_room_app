package shared_model

type Paging struct {
	Page       int `json:"page"`
	RowPerPage int `json:"row_per_page"`
	TotalRows  int `json:"total_rows"`
	TotalPages int `json:"total_pages"`
}
