package model

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type (
	User struct {
		ID           uuid.UUID
		Login        string
		PasswordHash string
		MasterHash   string
	}

	Secret struct {
		ID        uuid.UUID
		UserID    uuid.UUID
		DeviceID  uuid.UUID
		Data      string
		IsDeleted bool

		UploadedAt time.Time
		DeletedAt  time.Time
	}
)

func (u User) ValidateLogin(login string) error {
	if len(login) < 3 && len(login) > 60 {
		return errors.New("login not valid")
	}
	return nil
}
