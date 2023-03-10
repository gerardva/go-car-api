package models

type ErrorResponse struct {
	StatusCode   int
	Err  error
}

func (r *ErrorResponse) Error() string {
	return r.Err.Error()
}

func NewErrorResponse(statusCode int, err error) *ErrorResponse {
	return &ErrorResponse {
		StatusCode: statusCode,
		Err: err,
	}
}