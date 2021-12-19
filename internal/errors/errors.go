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

func (e ErrorResponse) Error() string {
	return e.Err.Error()
}

func (e ErrorResponse) ErrorType() ErrorType {
	return e.errorType
}

func NewError(err error, errtype ErrorType, statusCode int, msg string) error {
	return &ErrorResponse{
		statusCode: statusCode,
		Err:        err,
		Message:    msg,
		errorType:  errtype,
	}
}

// HTTP Errors
var (
	ErrUnauthorized   = ErrorType{"unauthorized"}
	ErrInvalidLogin   = ErrorType{"invalid-login"}
	ErrInvalidInput   = ErrorType{"invalid-input"}
	ErrDeleteData     = ErrorType{"delete-error"}
	ErrUpdateData     = ErrorType{"update-error"}
	ErrInternalServer = ErrorType{"internal-server-error"}
)

// Repository Errors
var (
	ErrUserNotFound = ErrorType{"user-not-found"}
)

func NewUnauthorizedError(err error) error {
	return NewError(err, ErrUnauthorized, 401, "unauthorized")
}

func NewInvalidLoginError(err error) error {
	return NewError(err, ErrInvalidLogin, 401, "invalid login")
}

func NewInvalidInputError(err error) error {
	return NewError(err, ErrInvalidInput, 422, "Data cannot be processed")
}

func NewDeleteDataError(err error) error {
	return NewError(err, ErrDeleteData, 400, "Data cannot be deleted")
}

func NewUpdateDataError(err error) error {
	return NewError(err, ErrDeleteData, 400, "Data cannot be updated")
}

func NewInternalServerError(err error) error {
	return NewError(err, ErrInternalServer, 400, "Internal server error")
}
