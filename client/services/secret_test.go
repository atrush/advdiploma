package services

import (
	"advdiploma/client/model"
	"advdiploma/client/pkg"
	"github.com/stretchr/testify/require"
	"testing"
)

func GetTestSecretSvc() *SecretService {
	cfg := &pkg.Config{
		MasterKey: "testKey",
	}

	svcSecret := NewSecret(cfg)
	return &svcSecret
}

func TestCard_ToSecret(t *testing.T) {
	secretSvc := GetTestSecretSvc()

	secret, err := secretSvc.ToSecret(model.TestCard.Info, model.TestCard)
	require.NoError(t, err)

	resObj, err := secretSvc.ReadFromSecret(secret)
	require.NoError(t, err)

	resCard, ok := resObj.(model.Card)
	require.True(t, ok)

	require.Equal(t, resCard, model.TestCard)
}
