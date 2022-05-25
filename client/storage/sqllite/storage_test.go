package sqllite

import (
	"advdiploma/client/model"
	"github.com/google/uuid"
	"github.com/icrowley/fake"
)

//var (
//	testSecret = model.Secret{
//		Info: model.Info{
//			TypeID:      model.SecretTypes["CARD"],
//			Title:       "Tinkoff Bank",
//			Description: "Tinka",
//			SecretID:    uuid.New(),
//			SecretVer:   1,
//			StatusID:    3,
//		},
//		SecretData: "T68WwfT8Kr1F3k21KBO8t1AqsALMW6A3xMt3BNKhKQWTOKtrRNKldTalvXt307jqax/C+Uag5so4PWlFVAeS6kM9jznhVSMR6n6in836UluABAtlxbZnCJX/i+WBIRhh4VVxjw3SaWo05/od5gYw5lzTgK8WNGMlbDPow==",
//	}
//)

func getMockSecret() model.Secret {
	return model.Secret{
		Info: model.Info{
			TypeID:      model.SecretTypes["CARD"],
			Title:       fake.Company(),
			Description: fake.CharactersN(200),
			SecretID:    uuid.New(),
			SecretVer:   1,
			StatusID:    3,
		},
		SecretData: fake.CharactersN(2000),
	}
}

func (s *TestSuite) DropSecretsTable() {
	_, err := s.storage.db.Exec("DELETE FROM Secrets")
	s.Require().NoError(err)
}

func (s *TestSuite) TestStorage_GetInfoList() {
	s.Run("Add list and get list of info", func() {
		count := 10
		arrSecrets := make([]model.Secret, count)

		for i := 0; i < count; i++ {
			secret := getMockSecret()

			id, err := s.storage.AddSecret(secret)
			s.Require().NoError(err)

			secret.ID = id

			arrSecrets[i] = secret
		}

		infoArr, err := s.storage.GetInfoList()
		s.Require().NoError(err)

		for i, el := range infoArr {
			s.Assert().EqualValues(arrSecrets[i].ID, el.ID)
			s.Assert().EqualValues(arrSecrets[i].TypeID, el.TypeID)
			s.Assert().EqualValues(arrSecrets[i].Title, el.Title)
			s.Assert().EqualValues(arrSecrets[i].Description, el.Description)
			s.Assert().EqualValues(arrSecrets[i].SecretID, el.SecretID)
			s.Assert().EqualValues(arrSecrets[i].SecretVer, el.SecretVer)
			s.Assert().EqualValues(arrSecrets[i].StatusID, el.StatusID)
		}
	})

	s.DropSecretsTable()
}

func (s *TestSuite) TestStorage_Add_Get() {
	s.Run("Add and read secret", func() {
		testSecret := getMockSecret()

		id, err := s.storage.AddSecret(testSecret)
		s.Require().NoError(err)

		dbSecret, err := s.storage.GetSecret(id)

		s.Assert().EqualValues(testSecret.TypeID, dbSecret.TypeID)
		s.Assert().EqualValues(testSecret.Title, dbSecret.Title)
		s.Assert().EqualValues(testSecret.Description, dbSecret.Description)
		s.Assert().EqualValues(testSecret.SecretID, dbSecret.SecretID)
		s.Assert().EqualValues(testSecret.SecretVer, dbSecret.SecretVer)
		s.Assert().EqualValues(testSecret.StatusID, dbSecret.StatusID)
		s.Assert().EqualValues(testSecret.SecretData, dbSecret.SecretData)
	})

	s.DropSecretsTable()
}
