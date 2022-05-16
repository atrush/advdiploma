package storage

import "advdiploma/client/model"

type Storage interface {
	AddSecret(v model.Secret, userID int) (int64, error)
	GetSecret(id int64) (model.Secret, error)
	GetInfoForUser(userID int) ([]model.Info, error)
	Close()
}
