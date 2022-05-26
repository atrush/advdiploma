package services

import (
	"advdiploma/client/model"
	"advdiploma/client/pkg"
	"advdiploma/client/storage"
	"encoding/json"
	"errors"
	"log"
)

type SecretService struct {
	cfg *pkg.Config
	db  storage.Storage
}

func NewSecret(cfg *pkg.Config, db storage.Storage) SecretService {
	return SecretService{
		cfg: cfg,
		db:  db,
	}
}

func (s *SecretService) NewSecret(secret model.Secret) (int64, error) {
	id, err := s.db.AddSecret(secret)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *SecretService) ToSecret(i interface{}) (model.Secret, error) {

	var info model.Info

	var data []byte
	var errMarshal error

	switch i.(type) {
	case model.Card:
		card, ok := i.(model.Card)
		if !ok {
			return model.Secret{}, errors.New("wrong Card type")
		}

		info = card.Info
		data, errMarshal = json.Marshal(card)

	case model.Auth:
		auth, ok := i.(model.Auth)
		if !ok {
			return model.Secret{}, errors.New("wrong Auth type")
		}

		info = auth.Info
		data, errMarshal = json.Marshal(auth)

	case model.Binary:
		bin, ok := i.(model.Binary)
		if !ok {
			return model.Secret{}, errors.New("wrong Binary type")
		}

		info = bin.Info
		data, errMarshal = json.Marshal(bin)

	default:
		return model.Secret{}, errors.New("wrong type")
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
		Info:       info,
		SecretData: encrypted,
	}, nil
}

func (s *SecretService) ReadFromSecret(el model.Secret) (interface{}, error) {

	decData, err := pkg.Decode(el.SecretData, s.cfg.MasterKey)
	if err != nil {
		log.Fatal("error:", err)
	}

	switch el.Info.TypeID {
	case model.SecretTypes["CARD"]:
		var card model.Card
		if err := json.Unmarshal(decData, &card); err != nil {
			return nil, errors.New("object is not Card type")
		}

		card.Info = el.Info

		return card, nil
	}

	return nil, errors.New("wrong TypeID")
}
