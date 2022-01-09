package response

type SuccessResponse struct {
	Data       interface{} `json:"data,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
	Code       int         `json:"code,omitempty"`
}

type ErrorResponse struct {
	DevMessage    string      `json:"dev_message,omitempty"`
	ClientMessage string      `json:"client_message,omitempty"`
	Code          int         `json:"code,omitempty"`
	Data          interface{} `json:"data,omitempty"`
}
