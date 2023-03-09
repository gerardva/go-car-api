package models

type ErrorResponse struct {
	StatusCode   int
	Error  error
}

func NewErrorResponse(statusCode int, err error) ErrorResponse {
	return ErrorResponse {
		StatusCode: statusCode,
		Error: err,
	}
}