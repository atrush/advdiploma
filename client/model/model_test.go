package model

import (
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestCard_ToSecret(t *testing.T) {
	d, err := TestCard.ToSecret()
	require.NoError(t, err)

	log.Printf("%+v", d)

	card := Card{}
	require.NoError(t, card.ReadFromSecret(d))

	require.Equal(t, TestCard, card)
}
