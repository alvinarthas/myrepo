package common

import "github.com/alvinarthas/myrepo/utils/logger"

type Error struct {
	Error   error  `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func RecordError(err error, code int, message string) Error {
	logger.Error(message, err)

	errorData := Error{
		Error:   err,
		Code:    code,
		Message: message,
	}

	return errorData
}
