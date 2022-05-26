package storage

import "advdiploma/client/model"

type Storage interface {
	AddSecret(v model.Secret) (int64, error)
	GetSecret(id int64) (model.Secret, error)
	GetMetaList() ([]model.Secret, error)

	//UpdateSecretBySecretID(v model.Secret) error
	UpdateSecret(v model.Secret) error
	Close()
}
