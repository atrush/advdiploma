package model

import (
	"errors"
	"github.com/google/uuid"
)

type (
	User struct {
		ID           uuid.UUID
		Login        string
		PasswordHash string
	}

	Secret struct {
		ID       int64
		UserID   int64
		ClientID int64
		Data     string
	}
)

func (u User) ValidateLogin(login string) error {
	if len(login) < 3 && len(login) > 60 {
		return errors.New("login not valid")
	}
	return nil
}
