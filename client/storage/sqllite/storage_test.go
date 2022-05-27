package sqllite

import (
	"advdiploma/client/model"
	"github.com/google/uuid"
	"github.com/icrowley/fake"
)

func (s *TestSuite) TestStorage_UpdateSecretBySecretID() {
	s.Run("Update exist", func() {

		secret1 := getMockSecret()

		id, err := s.storage.AddSecret(secret1)
		s.Require().NoError(err)

		secret1.ID = id

		//  mock updated secret
		secret2 := getMockSecret()
		secret2.ID = id
		secret2.SecretID = secret1.SecretID

		err = s.storage.UpdateSecretBySecretID(secret2)
		s.Require().NoError(err)

		dbRes, err := s.storage.GetSecret(id)
		s.Require().NoError(err)

		s.Assert().EqualValues(dbRes.ID, secret2.ID)
		s.Assert().EqualValues(dbRes.TypeID, secret2.TypeID)
		s.Assert().EqualValues(dbRes.Title, secret2.Title)
		s.Assert().EqualValues(dbRes.Description, secret2.Description)
		s.Assert().EqualValues(dbRes.SecretID, secret2.SecretID)
		s.Assert().EqualValues(dbRes.SecretVer, secret2.SecretVer)
		s.Assert().EqualValues(dbRes.StatusID, secret2.StatusID)

	})

	s.dropSecretsTable()

	s.Run("Update not exist", func() {
		secret1 := getMockSecret()
		secret1.ID = 201

		err := s.storage.UpdateSecretBySecretID(secret1)
		s.Require().Error(err)
	})

	s.dropSecretsTable()

	s.Run("Update secretID nil", func() {
		secret1 := getMockSecret()
		secret1.ID = 201
		secret1.SecretID = uuid.Nil

		err := s.storage.UpdateSecretBySecretID(secret1)
		s.Require().Error(err)
	})

	s.dropSecretsTable()
}

func (s *TestSuite) TestStorage_UpdateSecretByID() {
	s.Run("Update exist", func() {

		secret1 := getMockSecret()

		id, err := s.storage.AddSecret(secret1)
		s.Require().NoError(err)

		secret1.ID = id

		secret2 := getMockSecret()
		secret2.ID = id

		err = s.storage.UpdateSecret(secret2)
		s.Require().NoError(err)

		dbRes, err := s.storage.GetSecret(id)
		s.Require().NoError(err)

		s.Assert().EqualValues(dbRes.ID, secret2.ID)
		s.Assert().EqualValues(dbRes.TypeID, secret2.TypeID)
		s.Assert().EqualValues(dbRes.Title, secret2.Title)
		s.Assert().EqualValues(dbRes.Description, secret2.Description)
		s.Assert().EqualValues(dbRes.SecretID, secret2.SecretID)
		s.Assert().EqualValues(dbRes.SecretVer, secret2.SecretVer)
		s.Assert().EqualValues(dbRes.StatusID, secret2.StatusID)

	})

	s.dropSecretsTable()

	s.Run("Update not exist", func() {
		secret1 := getMockSecret()
		secret1.ID = 201

		err := s.storage.UpdateSecret(secret1)
		s.Require().Error(err)
	})

	s.dropSecretsTable()
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

		infoArr, err := s.storage.GetMetaList()
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

	s.dropSecretsTable()
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

	s.dropSecretsTable()
}

func (s *TestSuite) dropSecretsTable() {
	_, err := s.storage.db.Exec("DELETE FROM Secrets")
	s.Require().NoError(err)
}

func getMockSecret() model.Secret {
	return model.Secret{
		Info: model.Info{
			TypeID:      model.SecretTypes["CARD"],
			Title:       fake.Company(),
			Description: fake.CharactersN(200),
		},
		SecretID:   uuid.New(),
		SecretVer:  1,
		StatusID:   3,
		SecretData: fake.CharactersN(2000),
	}
}
