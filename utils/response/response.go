package response

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alvinarthas/myrepo/utils/common"
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

func Error(res http.ResponseWriter, errorData common.Error) {
	res.Header().Set("Content-Type", "application/json;charset=UTF-8")
	res.WriteHeader(errorData.Code)

	response := ErrorResponse{
		Caused:  errorData.Error.Error(),
		Message: errorData.Message,
		Code:    errorData.Code,
	}

	if err := json.NewEncoder(res).Encode(response); err != nil {
		log.Println(err)
	}
}
