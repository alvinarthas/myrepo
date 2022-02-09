package common

import logger "github.com/alvinarthas/myrepo/utils/log"

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
