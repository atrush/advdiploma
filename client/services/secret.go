package services

import (
	"advdiploma/client/model"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
)

func ToSecret(info model.Info, el interface{}) (model.Secret, error) {

	var data []byte
	var errMarshal error

	//  marshalling data
	switch info.TypeID {

	case model.SecretTypes["CARD"]:
		card, ok := el.(model.Card)
		if !ok {
			return model.Secret{}, errors.New("object is not Card type")
		}
		data, errMarshal = json.Marshal(card)

	default:
		return model.Secret{}, errors.New("wrong TypeID")
	}

	if errMarshal != nil {
		return model.Secret{}, errMarshal
	}

	//  encode data
	based64 := base64.StdEncoding.EncodeToString(data)

	return model.Secret{
		Info: info,
		Data: based64,
	}, nil
}

func ReadFromSecret(s model.Secret) (interface{}, error) {
	data, err := base64.StdEncoding.DecodeString(s.Data)
	if err != nil {
		log.Fatal("error:", err)
	}

	switch s.Info.TypeID {
	case model.SecretTypes["CARD"]:
		var card model.Card
		if err := json.Unmarshal(data, &card); err != nil {
			return nil, errors.New("object is not Card type")
		}

		card.Info = s.Info

		return card, nil
	}

	return nil, errors.New("wrong TypeID")
}
