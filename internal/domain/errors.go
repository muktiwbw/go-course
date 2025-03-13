package domain

import "errors"

var (
	// ErrPostNotFound is returned when a post is not found
	ErrPostNotFound = errors.New("post not found")
)
