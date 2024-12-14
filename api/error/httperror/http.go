package httperror

import "fmt"

type HttpError struct {
	ErrorCode   string `json:"errorCode"`
	Description string `json:"description"`
	Metadata    string `json:"metadata"`
	StatusCode  int    `json:"-"`
}

func (e HttpError) Error() string {
	return fmt.Sprintf("errorCode: %s, description: %s,  metadata: %s", e.ErrorCode, e.Description, e.Metadata)
}

func NewWithMetadata(key, metadata string) HttpError {
	return NewWithStatus(key, metadata, 0)
}

func NewWithStatus(key, metadata string, status int) HttpError {
	if err, ok := errors[key]; ok {
		err.ErrorCode = key
		err.Metadata = metadata
		if status != 0 {
			err.StatusCode = status
		}
		return err
	}
	return errors[UndefinedErrorCode]
}
