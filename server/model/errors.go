package model

import "errors"

var (
	ErrorConflictSaveUser = errors.New("user already exist")
	ErrorItemNotFound     = errors.New("item not found")

	ErrorWrongAuthData = errors.New("incorrect login or password")
)
