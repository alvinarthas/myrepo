package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func Success(res http.ResponseWriter, httpCode int, data interface{}, pagination interface{}) {
	res.Header().Set("Content-Type", "application/json;charset=UTF-8")
	res.WriteHeader(httpCode)

	response := SuccessResponse{
		Code:       httpCode,
		Data:       data,
		Pagination: pagination,
	}

	if err := json.NewEncoder(res).Encode(response); err != nil {
		// logger.Fatal(err.Error(), nil)
		log.Println(err)
	}
}
