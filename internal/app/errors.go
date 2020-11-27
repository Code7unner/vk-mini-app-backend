package app

import "errors"

var (
	// ErrUserNotFound user with given id not found.
	ErrUserNotFound = errors.New("user not found")
)
