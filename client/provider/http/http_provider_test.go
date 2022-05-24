package http

import (
	"advdiploma/client/provider/http/model"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"net/http"
	"testing"
	"time"
)

var (
	token = fake.CharactersN(16)

	authData = model.LoginRequest{
		Login:      fake.CharactersN(8),
		Password:   fake.CharactersN(8),
		MasterHash: fake.CharactersN(16),
		DeviceID:   uuid.New(),
	}

	srvBaseCfg = serverTestConfig{
		returnHeaders: map[string]string{"Authorization": token},
		returnBody:    "",
		returnStatus:  http.StatusOK,
		sleep:         0,

		reqHeaders: map[string]string{"content-type": "application/json"},
		reqBody:    mustMarshalLoginRequest(authData),
	}

	provBaseCfg = HTTPConfig{
		AuthURL:     "/api/user/login",
		RegisterURL: "/api/user/register",
		timeout:     time.Millisecond * 200,
	}
)

func TestProvider_Auth(t *testing.T) {
	authSrvConfig := srvBaseCfg.New(withReqURL(provBaseCfg.AuthURL))

	tests := []struct {
		name      string
		serverCfg serverTestConfig

		reqErr   assert.ErrorAssertionFunc
		reqToken string
		wait     time.Duration
	}{
		{
			name:      "register ok",
			serverCfg: authSrvConfig,
			reqErr:    assert.NoError,
			reqToken:  token,
		},
		{
			name:      "too long request",
			serverCfg: authSrvConfig.New(withSleep(provBaseCfg.timeout + time.Millisecond*200)),
			reqErr:    assert.Error,
		},
		{
			name:      "err 400",
			serverCfg: authSrvConfig.New(withReturnStatus(http.StatusBadRequest)),
			reqErr:    assert.Error,
		},
		{
			name:      "err 401",
			serverCfg: authSrvConfig.New(withReturnStatus(http.StatusUnauthorized)),
			reqErr:    assert.Error,
		},
		{
			name:      "err 500",
			serverCfg: authSrvConfig.New(withReturnStatus(http.StatusInternalServerError)),
			reqErr:    assert.Error,
		},
		{
			name:      "err 415",
			serverCfg: authSrvConfig.New(withReturnStatus(http.StatusUnsupportedMediaType)),
			reqErr:    assert.Error,
		},
		{
			name:      "err 403",
			serverCfg: authSrvConfig.New(withReturnStatus(http.StatusForbidden)),
			reqErr:    assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			server := getTestHTTPServer(t, tt.serverCfg)
			defer server.Close()

			provCfg := provBaseCfg
			provCfg.BaseURL = server.URL

			provider := NewHTTPProvider(provCfg)
			err := provider.Authorise(authData.Login, authData.Password, authData.MasterHash, authData.DeviceID)

			tt.reqErr(t, err)

			//  check client authorised
			require.EqualValues(t, *provider.client.apiToken, tt.reqToken)
		})
	}
}

func TestProvider_Register(t *testing.T) {
	regSrvConfig := srvBaseCfg.New(withReqURL(provBaseCfg.RegisterURL))

	tests := []struct {
		name      string
		serverCfg serverTestConfig

		reqErr   assert.ErrorAssertionFunc
		reqToken string
		wait     time.Duration
	}{
		{
			name:      "register ok",
			serverCfg: regSrvConfig,
			reqErr:    assert.NoError,
			reqToken:  token,
		},
		{
			name:      "too long request",
			serverCfg: regSrvConfig.New(withSleep(provBaseCfg.timeout + time.Millisecond*200)),
			reqErr:    assert.Error,
		},
		{
			name:      "err 400",
			serverCfg: regSrvConfig.New(withReturnStatus(http.StatusBadRequest)),
			reqErr:    assert.Error,
		},
		{
			name:      "err 401",
			serverCfg: regSrvConfig.New(withReturnStatus(http.StatusUnauthorized)),
			reqErr:    assert.Error,
		},
		{
			name:      "err 500",
			serverCfg: regSrvConfig.New(withReturnStatus(http.StatusInternalServerError)),
			reqErr:    assert.Error,
		},
		{
			name:      "err 415",
			serverCfg: regSrvConfig.New(withReturnStatus(http.StatusUnsupportedMediaType)),
			reqErr:    assert.Error,
		},
		{
			name:      "err 403",
			serverCfg: regSrvConfig.New(withReturnStatus(http.StatusForbidden)),
			reqErr:    assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			server := getTestHTTPServer(t, tt.serverCfg)
			defer server.Close()

			provCfg := provBaseCfg
			provCfg.BaseURL = server.URL

			provider := NewHTTPProvider(provCfg)
			err := provider.Register(authData.Login, authData.Password, authData.MasterHash, authData.DeviceID)

			tt.reqErr(t, err)

			//  check client authorised
			require.EqualValues(t, *provider.client.apiToken, tt.reqToken)
		})
	}
}

func mustMarshalLoginRequest(val model.LoginRequest) string {
	res, err := json.Marshal(val)
	if err != nil {
		log.Fatal(err)
	}
	return string(res)
}
