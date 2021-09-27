package services

import (
	"errors"
)

// TODO FIND A BETTER LOCATION TO ERRORS
var (
	ErrUserIDMustBePositive  = errors.New("user_id must be greater than 0")
	ErrUserIDMustBeValid     = errors.New("user_id must be valid")
	ErrInvalidUserOrPassword = errors.New("invalid user or password")
	ErrUserDoesntExists      = errors.New("user doesn't exists")

	ErrThereIsNoFiles = errors.New("there is no files")
)
