package auth

import (
	"advdiploma/server/model"
	"context"
	"github.com/google/uuid"
)

//  Authenticator is the interface that wraps methods user identification, authentication, authorisation.
type Authenticator interface {
	CreateUser(ctx context.Context, login string, password string) (model.User, error)
	Authenticate(ctx context.Context, login string, password string) (model.User, error)
	EncodeTokenUserID(userID uuid.UUID, deviceID uuid.UUID) (string, error)
}