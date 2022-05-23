package model

import (
	"fmt"
	"github.com/google/uuid"
)

type (
	ContextKey string

	LoginRequest struct {
		Login      string    `json:"login"`
		MasterHash string    `json:"master_hash"`
		Password   string    `json:"password"`
		DeviceID   uuid.UUID `json:"device_id"`
	}

	UserContextData struct {
		UserID   uuid.UUID
		DeviceID uuid.UUID
	}
)

func (r LoginRequest) Validate() error {
	if len(r.Login) < 3 {
		return fmt.Errorf("login must be larger then 3 symbols")
	}
	if len(r.Login) > 60 {
		return fmt.Errorf("login must be less then 60 symbols")
	}
	if len(r.Password) < 3 {
		return fmt.Errorf("password must be larger then 3 symbols")
	}
	if len(r.MasterHash) < 3 {
		return fmt.Errorf("master hash must be larger then 3 symbols")
	}
	if r.DeviceID == uuid.Nil {
		return fmt.Errorf("device id is empty")
	}
	return nil
}

//
//func (s *SecretRequst) ToCanonical(userID uuid.UUID) (model.Secret, error) {
//	//todo validate
//	return model.Secret{
//		Data:   s.Data,
//		UserID: userID,
//	}, nil
//}

var (
	ContextKeyUserID = ContextKey("user-id")
)
