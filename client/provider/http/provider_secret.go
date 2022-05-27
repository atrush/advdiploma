package http

import (
	"advdiploma/client/provider/http/model"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
)

func (p *HTTPProvider) Upload(data string, id uuid.UUID, ver int) (uuid.UUID, int, error) {
	secretData, err := json.Marshal(model.SecretRequest{
		Data: data,
		ID:   id,
		Ver:  ver,
	})
	if err != nil {
		return uuid.Nil, 0, fmt.Errorf("upload error: %w", err)
	}

	//  prepare request
	request, err := http.NewRequest(http.MethodPost, p.cfg.SecretURL, bytes.NewBuffer(secretData))
	if err != nil {
		return uuid.Nil, 0, fmt.Errorf("request error: %w", err)
	}

	request.Header.Set("Content-Type", "Content-Type: text/plain")

	//  do request
	response, err := p.client.Do(request)
	if err != nil {
		return uuid.Nil, 0, fmt.Errorf("request error: %w", err)
	}

	//  read body
	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return uuid.Nil, 0, fmt.Errorf("request error: %w", err)
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Println(err.Error())
		}
	}()

	var respSecret model.SecretRequest
	if err := json.Unmarshal(respBody, &respSecret); err != nil {
		return uuid.Nil, 0, fmt.Errorf("request error: %w", err)
	}

	return respSecret.ID, respSecret.Ver, nil
}
