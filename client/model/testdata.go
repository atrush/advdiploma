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
	TestSecret = Secret{
		Info: Info{
			TypeID:      SecretTypes["CARD"],
			Title:       "Tinkoff Bank",
			Description: "Tinka",
		},
		Data: "T68WwfT8Kr1F3k21KBO8t1AqsALMW6A3xMt3BNKhKQWTOKtrRNKldTalvXt307jqax/C+Uag5so4PWlFVAeS6kM9jznhVSMR6n6in836UluABAtlxbZnCJX/i+WBIRhh4VVxjw3SaWo05/od5gYw5lzTgK8WNGMlbDPow==",
	}
)
