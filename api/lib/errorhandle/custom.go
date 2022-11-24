package errorhandle

import "errors"

type CustomError struct {
	Status int
	Err error
}

// ここを経由することで CustomError を error に偽装する
func InitCustomError(status int, err error) error {
	return CustomError{status, err}
}

// CustomError.Error() を使えるようにすることで
// CustomError を error に偽装する
func (ce CustomError) Error() string {
	if ce.Err != nil {
		return ce.Err.Error()
	}
	return "unknown error"
}

func (ce CustomError) StatusCode() int {
	return ce.Status
}

func (ce CustomError) Unwrap() error {
    return ce.Err // Returns inner error
}

// response 用の parse
type ParsedError struct {
	Status int
	Message string
	OrgErr string
}

func InitParsedError(err error) ParsedError {
	var ce CustomError
	status := 400
	unwrappedErr := err
	if errors.As(err, &ce) {
		status = ce.Status
		unwrappedErr = ce.Unwrap()
	}

	return ParsedError{
		Status: status,
		Message: unwrappedErr.Error(),
		OrgErr: err.Error(),
	}
}