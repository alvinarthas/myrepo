package model

type Dummy struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Adress   string `json:"address"`
	IsActive bool   `json:"is_active"`
	Type     string `json:"type"`
}
