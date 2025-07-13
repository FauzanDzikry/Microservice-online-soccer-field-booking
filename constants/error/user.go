package error

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrPasswordIncorrect = errors.New("password incorrect")
	ErrPasswordNotMatch  = errors.New("password not match")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotActive     = errors.New("user not active")
)

var UserErrors = []error{
	ErrUserNotFound,
	ErrUserAlreadyExists,
	ErrUserNotActive,
	ErrPasswordIncorrect,
	ErrPasswordNotMatch,
}
