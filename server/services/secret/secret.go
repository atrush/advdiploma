package secret

import (
	"advdiploma/server/model"
	"advdiploma/server/storage"
	"context"
	"github.com/google/uuid"
)

var _ SecretManager = (*Secret)(nil)

type Secret struct {
	storage storage.Storage
}

func NewSecret(s storage.Storage) (*Secret, error) {
	return &Secret{
		storage: s,
	}, nil
}

func (s *Secret) Add(ctx context.Context, secret model.Secret) (model.Secret, error) {
	return s.storage.Secret().Add(ctx, secret)
}
func (s *Secret) Get(ctx context.Context, id uuid.UUID) (model.Secret, error) {
	return s.storage.Secret().Get(ctx, id)
}
