package model

import "github.com/google/uuid"

type (
	ContextKey string
)

type UserContextData struct {
	UserID   uuid.UUID
	DeviceID uuid.UUID
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
