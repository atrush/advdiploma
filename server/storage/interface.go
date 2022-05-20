package storage

import (
	"advdiploma/server/model"
	"context"
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
	Add(ctx context.Context, secret model.Secret) (model.Secret, error)
	Get(ctx context.Context, id int64) (model.Secret, error)
}
