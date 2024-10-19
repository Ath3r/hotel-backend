package constants

import "errors"

var (
	ErrUnexpected = errors.New("unexpected error")
	ErrNotFound = errors.New("not found")
	ErrBadRequest = errors.New("bad request")
	ErrInternalServer = errors.New("internal server error")

	// config
	ErrConfigLoad = errors.New("failed to load config")
	ErrParseConfig = errors.New("failed to parse config")
	ErrEmptyVar    = errors.New("required variable is empty")
)