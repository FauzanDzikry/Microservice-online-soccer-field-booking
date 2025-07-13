package error

import "errors"

var (
	ErrInternalServer  = errors.New("internal server error")
	ErrSQL             = errors.New("database server failed to execute query")
	ErrTooManyRequests = errors.New("too many requests")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrInvalidToken    = errors.New("invalid token")
	ErrForbidden       = errors.New("forbidden")
	ErrNotFound        = errors.New("not found")
)

var GeneralErrors = []error{
	ErrInternalServer,
	ErrSQL,
	ErrTooManyRequests,
	ErrUnauthorized,
	ErrInvalidToken,
	ErrForbidden,
	ErrNotFound,
}
