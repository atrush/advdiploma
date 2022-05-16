package services

import (
	"advdiploma/client/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCard_ToSecret(t *testing.T) {

	secret, err := ToSecret(model.TestCard.Info, model.TestCard)
	require.NoError(t, err)

	resObj, err := ReadFromSecret(secret)
	require.NoError(t, err)

	resCard, ok := resObj.(model.Card)
	require.True(t, ok)

	require.Equal(t, resCard, model.TestCard)
}
