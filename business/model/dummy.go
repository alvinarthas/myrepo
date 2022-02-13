package model

type Dummy struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	IsActive bool   `json:"is_active"`
	Type     string `json:"type"`
}

// Request

type CreateDummyRequest struct {
	Name    string `json:"name" validate:"required"`
	Age     int    `json:"age" validate:"required"`
	Address string `json:"address"`
	Type    string `json:"type" oneof:"Human Mutant Undead"`
}

type UpdateDummyRequest struct {
	ID      string `json:"id" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Age     int    `json:"age" validate:"required"`
	Address string `json:"address"`
	Type    string `json:"type" oneof:"Human Mutant Undead"`
}

// Response

type GetDummyByIDResponse struct {
	Data Dummy `json:"data"`
}

type GetDummyListResponse struct {
	Data       []Dummy    `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type CreateDummyResponse struct {
	Data []string `json:"data"`
}

type UpdateDummyResponse struct {
	Data []string `json:"data"`
}

type DeleteDummyResponse struct {
	Data []string `json:"data"`
}
