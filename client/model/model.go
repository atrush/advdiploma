package model

import (
	"encoding/base64"
	"encoding/json"
	"log"
)

var (
	SecretTypes = map[string]int{
		"CARD":   1,
		"AUTH":   2,
		"TEXT":   3,
		"BINARY": 4,
	}
)

type Info struct {
	ID          int64
	TypeID      int
	Title       string
	Description string
}

//Cardholder name
//PAN (Primary Account Number) (the 16 digit number on the front of the card)
//Expiration date
//Service code (You won’t find this data on the card itself. It lives within the magnetic stripe
type Card struct {
	Info            Info   `json:"-"`
	CardholderName  string `json:"cardholder"`
	PAN             string `json:"pan"`
	ExpirationMonth int    `json:"expiration_month"`
	ExpirationYear  int    `json:"expiration_year"`
	ServiceCode     int    `json:"code"`
}

func (c *Card) ToSecret() (Secret, error) {

	js, err := json.Marshal(c)
	if err != nil {
		return Secret{}, err
	}
	based64 := base64.StdEncoding.EncodeToString(js)

	s := Secret{
		Info: c.Info,
		Data: based64,
	}
	return s, nil
}

func (c *Card) ReadFromSecret(s Secret) error {
	data, err := base64.StdEncoding.DecodeString(s.Data)
	if err != nil {
		log.Fatal("error:", err)
	}

	if err := json.Unmarshal(data, c); err != nil {
		return err
	}

	c.Info = s.Info

	return nil
}

type Auth struct {
	Info     Info   `json:"-"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Binary struct {
	Info        Info `json:"-"`
	Data        []byte
	ContentType string
	Filename    string
}

type Secret struct {
	Info Info
	Data string
}
