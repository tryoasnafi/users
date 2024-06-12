package users

import "errors"

type RepositoryError error

var (
	ErrUserNotFound RepositoryError = errors.New("user not found")
)
