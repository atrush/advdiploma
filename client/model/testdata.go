package model

var (
	TestCard = Card{
		Info: Info{
			TypeID:      SecretTypes["CARD"],
			Title:       "Tinkoff Bank",
			Description: "Tinka",
		},
		CardholderName:  "PETR IVANOV",
		ExpirationMonth: 2,
		ExpirationYear:  25,
		PAN:             "111 2222 5525 4544",
		ServiceCode:     546,
	}
)
