package services

import (
	"advdiploma/client/model"
	"advdiploma/client/pkg"
	"advdiploma/client/storage"
	mk "advdiploma/client/storage/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func GetTestSecretSvc(t *testing.T, storage storage.Storage) *SecretService {
	cfg := &pkg.Config{
		MasterKey: "testKey",
	}

	svcSecret := NewSecret(cfg, storage)
	return &svcSecret
}

func TestCard_ToSecret(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		obj     interface{}
		storage storage.Storage
		reqErr  assert.ErrorAssertionFunc
	}{
		{
			name:    "card",
			obj:     model.TestCard,
			storage: storageEmpty(ctrl),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			secretSvc := GetTestSecretSvc(t, tt.storage)

			secret, err := secretSvc.ToSecret(tt.obj)
			require.NoError(t, err)

			resObj, err := secretSvc.ReadFromSecret(secret)
			require.NoError(t, err)

			switch resObj.(type) {
			case model.Card:
				require.Equal(t, resObj.(model.Card), tt.obj)

			default:
				t.Error("wrong type")
			}

		})
	}
}

func storageEmpty(ctrl *gomock.Controller) *mk.MockStorage {
	storageMock := mk.NewMockStorage(ctrl)
	return storageMock
}
