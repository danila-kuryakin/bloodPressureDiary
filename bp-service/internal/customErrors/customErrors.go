package customErrors

import "errors"

var (
	ErrInvalidTagName    = errors.New("invalid tag name")
	ErrTagNotOwnedByUser = errors.New("tag does not belong to user")

	ErrMetadataNotProvided            = errors.New("metadata not provided")
	ErrTokenApiServiceIsNotEquivalent = errors.New("The token API service is not equivalent")
)
