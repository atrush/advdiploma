package http

import (
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTokenClient_SetToken(t *testing.T) {
	client := NewTokenClient()
	testTokenClient(t, client, "")

	token := fake.CharactersN(32)
	client.SetToken(token)
	testTokenClient(t, client, token)
}

func testTokenClient(t *testing.T, client *TokenClient, reqToken string) {
	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			token := req.Header.Get("Authorization")

			if _, err := rw.Write([]byte(token)); err != nil {
				log.Fatal(err.Error())
			}
		}))
	defer server.Close()

	resp, err := client.Get(server.URL)
	require.NoError(t, err)

	body, err := ioutil.ReadAll(resp.Body)
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	require.EqualValues(t, body, reqToken)
}
