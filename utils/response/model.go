package response

type SuccessResponse struct {
	Data       interface{} `json:"data,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
	Code       int         `json:"code,omitempty"`
}

type ErrorResponse struct {
	Caused  string `json:"caused,omitempty"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}
