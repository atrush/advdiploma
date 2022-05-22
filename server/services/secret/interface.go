package secret

import (
	"advdiploma/server/model"
	"context"
	"github.com/google/uuid"
)

type SecretManager interface {
	Add(ctx context.Context, secret model.Secret) (model.Secret, error)
	Get(ctx context.Context, id uuid.UUID) (model.Secret, error)
}
