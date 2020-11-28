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

var (
	// ErrSteamUserNotFound steam user with given id not found.
	ErrSteamUserNotFound = errors.New("steam user not found")
)

var (
	// ErrMatchNotFound team with given id not found.
	ErrMatchNotFound = errors.New("match not found")
)