package model

type Page struct {
	PageSize      int `json:"page_size"`
	CurrentPage   int `json:"current_page"`
	TotalCount    int `json:"total_count"`
	TotalPagesNum int `json:"total_pages_num"`
}