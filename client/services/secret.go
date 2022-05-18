package services

import (
	"advdiploma/client/model"
	"advdiploma/client/pkg"
	"encoding/json"
	"errors"
	"log"
)

type SecretService struct {
	cfg *pkg.Config
}

func NewSecret(cfg *pkg.Config) SecretService {
	return SecretService{
		cfg: cfg,
	}
}

func (s *SecretService) ToSecret(info model.Info, el interface{}) (model.Secret, error) {

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
	encrypted, err := pkg.Encode(data, s.cfg.MasterKey)
	if err != nil {
		return model.Secret{}, err
	}

	return model.Secret{
		Info: info,
		Data: encrypted,
	}, nil
}

func (s *SecretService) ReadFromSecret(el model.Secret) (interface{}, error) {

	data, err := pkg.Decode(el.Data, s.cfg.MasterKey)
	if err != nil {
		log.Fatal("error:", err)
	}

	switch el.Info.TypeID {
	case model.SecretTypes["CARD"]:
		var card model.Card
		if err := json.Unmarshal(data, &card); err != nil {
			return nil, errors.New("object is not Card type")
		}

		card.Info = el.Info

		return card, nil
	}

	return nil, errors.New("wrong TypeID")
}
