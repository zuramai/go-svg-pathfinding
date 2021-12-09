package errors

type ErrorType struct {
	t string
}
type ErrorResponse struct {
	statusCode int       `json:"-"`
	Message    string    `json:"message"`
	Err        error     `json:"-"`
	errorType  ErrorType `json:"-"`
}

func (e ErrorResponse) Error() error {
	return e.Err
}

func (e ErrorResponse) ErrorType() ErrorType {
	return e.errorType
}

var (
	UnauthorizedError = ErrorType{"unauthorized"}
	InvalidLoginError = ErrorType{"invalid-login"}
	InvalidInputError = ErrorType{"invalid-input"}
	DeleteDataError   = ErrorType{"delete-error"}
)

func NewErrorResponse(err error, statusCode int, msg string) *ErrorResponse {
	return &ErrorResponse{
		statusCode: statusCode,
		Err:        err,
		Message:    msg,
	}
}
