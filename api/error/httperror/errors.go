package httperror

import (
	"net/http"
)

const (
	UndefinedErrorCode = "ERR-0000"
	InvalidRequest     = "ERR-0001"
)

var errors = map[string]HttpError{
	InvalidRequest: {
		StatusCode:  http.StatusBadRequest,
		Description: "Request is invalid. Please provide valid request",
	},
}
