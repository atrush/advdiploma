package model

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

type Informer interface {
	GetInfo() Info
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

func (c *Card) GetInfo() Info {
	return c.Info
}

var _ Informer = (*Card)(nil)

type Auth struct {
	Info     Info   `json:"-"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (a *Auth) GetInfo() Info {
	return a.Info
}

type Binary struct {
	Info        Info `json:"-"`
	Data        []byte
	ContentType string
	Filename    string
}

func (b *Binary) GetInfo() Info {
	return b.Info
}

type Secret struct {
	Info Info
	Data string
}
