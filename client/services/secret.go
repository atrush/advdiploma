package services

import (
	"advdiploma/client/model"
	"advdiploma/client/pkg"
	"advdiploma/client/storage"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
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

func (s *SecretService) AddAuth(el model.Auth) (int64, error) {
	return s.addSecret(el)
}

func (s *SecretService) AddCard(el model.Card) (int64, error) {
	return s.addSecret(el)
}

func (s SecretService) ReadBinary(filePath string) (model.Binary, error) {
	b := model.Binary{
		Info: model.Info{
			TypeID: model.SecretTypes["BINARY"],
		},
	}

	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err.Error())
	}

	b.Filename = filepath.Base(filePath)
	b.ContentType = http.DetectContentType(bytes)
	b.Data = bytes

	return b, nil
}

func (s *SecretService) AddBinary(filePath string, title string, description string) (int64, error) {
	b, err := s.ReadBinary(filePath)
	if err != nil {
		return 0, nil
	}

	b.Title = title
	b.Description = description

	return s.addSecret(b)
}

func (s *SecretService) addSecret(obj interface{}) (int64, error) {
	secret, err := s.ToSecret(obj)
	if err != nil {
		return 0, err
	}

	secret.StatusID = model.SecretStatuses["NEW"]
	secret.SecretID = uuid.Nil
	secret.SecretVer = 1

	id, err := s.db.AddSecret(secret)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *SecretService) UpdateSecret(secret model.Secret) error {
	//  if el secret id == nil, el not uploaded to server, must stay status NEW
	if secret.SecretID != uuid.Nil {
		secret.StatusID = model.SecretStatuses["EDITED"]
	}

	if err := s.db.UpdateSecret(secret); err != nil {
		return err
	}

	return nil
}

func (s *SecretService) GetSecret(id int64) (model.Secret, error) {
	dbSecret, err := s.db.GetSecret(id)
	if err != nil {
		return model.Secret{}, err
	}
	return dbSecret, nil
}
func (s *SecretService) GetSecretBySecretID(id uuid.UUID) (model.Secret, error) {
	dbSecret, err := s.db.GetSecretByExtID(id)
	if err != nil {
		return model.Secret{}, err
	}
	return dbSecret, nil
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

	case model.Text:
		txt, ok := i.(model.Text)
		if !ok {
			return model.Secret{}, errors.New("wrong Text type")
		}

		info = txt.Info
		data, errMarshal = json.Marshal(txt)

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

	case model.SecretTypes["AUTH"]:
		var auth model.Auth
		if err := json.Unmarshal(decData, &auth); err != nil {
			return nil, errors.New("object is not Auth type")
		}

		auth.Info = el.Info

		return auth, nil

	case model.SecretTypes["TEXT"]:
		var txt model.Text
		if err := json.Unmarshal(decData, &txt); err != nil {
			return nil, errors.New("object is not Text type")
		}

		txt.Info = el.Info

		return txt, nil

	case model.SecretTypes["BINARY"]:
		var bn model.Binary
		if err := json.Unmarshal(decData, &bn); err != nil {
			return nil, errors.New("object is not Binary type")
		}

		bn.Info = el.Info

		return bn, nil
	}

	return nil, errors.New("wrong TypeID")
}

func (s *SecretService) DeleteSoftSecret(id int64) error {
	dbSecret, err := s.db.GetSecret(id)
	if err != nil {
		//  if not found ok
		if errors.Is(err, model.ErrorItemNotFound) {
			return nil
		}

		return err
	}

	dbSecret.StatusID = model.SecretStatuses["DELETED"]
	if err := s.db.UpdateSecret(dbSecret); err != nil {

		return err
	}

	return nil
}
