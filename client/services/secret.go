package services

import (
	"advdiploma/client/model"
	"advdiploma/client/pkg"
	"advdiploma/client/storage"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
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

// new -> upload
// update ->download
// changed -> toupload, todownload - ask collision
// delete -> if new, delete if actual mark to delete -  post delete, deleted+ver

func (s *SecretService) AddSecret(obj interface{}) (int64, error) {
	secret, err := s.ToSecret(obj)
	if err != nil {
		return 0, err
	}

	secret.StatusID = model.SecretStatuses["NEW"]
	secret.SecretID = uuid.Nil
	secret.SecretVer = 0

	id, err := s.db.AddSecret(secret)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *SecretService) UpdateSecret(id int64, obj interface{}) error {
	secret, err := s.ToSecret(obj)
	if err != nil {
		return err
	}

	dbSecret, err := s.db.GetSecret(id)
	if err != nil {
		return err
	}

	dbSecret.Info = secret.Info
	dbSecret.SecretData = secret.SecretData
	dbSecret.StatusID = model.SecretStatuses["UPDATED"]

	err = s.db.UpdateSecret(dbSecret)
	if err != nil {
		return err
	}

	return nil
}

func (s *SecretService) DeleteSecret(id int64) error {
	dbSecret, err := s.db.GetSecret(id)
	if err != nil {
		return err
	}

	dbSecret.StatusID = model.SecretStatuses["DELETED"]

	err = s.db.UpdateSecret(dbSecret)
	if err != nil {
		return err
	}

	return nil
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

func (s *SecretService) ReadInfoFromSecret(enc string) (model.Info, error) {
	decData, err := pkg.Decode(enc, s.cfg.MasterKey)
	if err != nil {
		log.Fatal("error:", err)
	}
	var info model.Info

	err = json.Unmarshal(decData, &info)

	if err != nil {
		log.Println(err.Error())
	}

	return info, nil
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
