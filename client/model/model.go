package model

import "github.com/google/uuid"

var (
	SecretTypes = map[string]int{
		"CARD":   1,
		"AUTH":   2,
		"TEXT":   3,
		"BINARY": 4,
	}

	SecretStatuses = map[string]int{
		"NEW":      1,
		"UPLOAD":   2,
		"DOWNLOAD": 3,
	}
)

type Info struct {
	TypeID      int    `json:"type_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Informer interface {
	GetInfo() Info
}

//Cardholder name
//PAN (Primary Account Number) (the 16 digit number on the front of the card)
//Expiration date
//Service code (You won’t find this data on the card itself. It lives within the magnetic stripe
type Card struct {
	Info
	CardholderName  string `json:"cardholder"`
	PAN             string `json:"pan"`
	ExpirationMonth int    `json:"expiration_month"`
	ExpirationYear  int    `json:"expiration_year"`
	ServiceCode     int    `json:"code"`
}

func (c *Card) GetInfo() Info {
	return c.Info
}

var _ Informer = (*Card)(nil)

type Auth struct {
	Info
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (a *Auth) GetInfo() Info {
	return a.Info
}

type Binary struct {
	Info
	Data        []byte
	ContentType string
	Filename    string
}

type Secret struct {
	Info

	ID        int64
	SecretID  uuid.UUID
	SecretVer int
	StatusID  int

	SecretData string
}
