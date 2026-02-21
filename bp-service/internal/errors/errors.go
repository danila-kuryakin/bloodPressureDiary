package errors

import "errors"

var (
	ErrInvalidTagName    = errors.New("invalid tag name")
	ErrTagNotOwnedByUser = errors.New("tag does not belong to user")
)
