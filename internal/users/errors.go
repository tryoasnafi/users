package users

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrNeedUserID = errors.New("need user id parameter")
)