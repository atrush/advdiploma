package storage

import (
	"advdiploma/server/model"
	"context"
	"github.com/google/uuid"
)

type Storage interface {
	//  User returns repository for working with users.
	User() UserRepository
	//  Order returns repository for working with orders.
	Secret() SecretRepository
}

type UserRepository interface {
	//  Adds new user to storage
	Create(ctx context.Context, user model.User) (model.User, error)
	//  Returns user from storage
	GetByLogin(ctx context.Context, login string) (model.User, error)
}

type SecretRepository interface {
	Add(ctx context.Context, secret model.Secret) (uuid.UUID, error)
	Get(ctx context.Context, id uuid.UUID) (model.Secret, error)
	GetUserVersionList(ctx context.Context, userID uuid.UUID) (map[uuid.UUID]int, error)
}
