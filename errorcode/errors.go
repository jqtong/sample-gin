package errorcode

// APIError struct
type APIError struct {
	Code int
	Msg  string
	Err  *error
}

// New returns new APIError struct
func New(code int, msg string, err error) *APIError {

	return &APIError{
		Code: code,
		Msg:  msg,
		Err:  &err,
	}
}

func (e *APIError) Error() string {
	return ""
}
