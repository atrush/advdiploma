package sqllite

import "advdiploma/client/model"

func (s *TestSuite) TestStorage_Add_Get() {
	s.Run("Add and read secret", func() {
		secret, err := model.TestCard.ToSecret()
		s.Require().NoError(err)

		id, err := s.storage.AddSecret(secret, 2)
		s.Require().NoError(err)

		dbSecret, err := s.storage.GetSecret(id)

		s.Require().Equal(secret.Info.TypeID, dbSecret.Info.TypeID)
		s.Require().Equal(secret.Info.Title, dbSecret.Info.Title)
		s.Require().Equal(secret.Info.Description, dbSecret.Info.Description)
		s.Require().Equal(secret.Data, dbSecret.Data)
	})
}