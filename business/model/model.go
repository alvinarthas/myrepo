package model

type Pagination struct {
	Page       int `json:"page,omitnull"`
	DataInPage int `json:"data_in_page,omitnull"`
	TotalData  int `json:"total_data,omitnull"`
	TotalPage  int `json:"total_page,omitnull"`
}
