package app

import "errors"

var (
	// ErrUserNotFound user with given id not found.
	ErrUserNotFound = errors.New("user not found")
)

var (
	// ErrTeamNotFound team with given id not found.
	ErrTeamNotFound = errors.New("team not found")
)