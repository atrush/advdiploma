package provider

import (
	"github.com/google/uuid"
)

type SecretProvider interface {
	Authorise(login string, pass string, masterHash string, deviceID uuid.UUID) error
	Register(login string, pass string, masterHash string, deviceID uuid.UUID) error

	GetSyncList() (map[uuid.UUID]int, error)
	//GetSecret(id uuid.UUID) (model.Secret, error)
	//UploadSecret(secret model.Secret) (uuid.UUID, error)
}
