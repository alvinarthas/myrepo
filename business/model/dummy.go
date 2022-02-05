package model

type Dummy struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	IsActive bool   `json:"is_active"`
	Type     string `json:"type"`
}
