package sqllite

import (
	"advdiploma/client/model"
)

func (s *TestSuite) TestStorage_Add_Get() {
	s.Run("Add and read secret", func() {

		id, err := s.storage.AddSecret(model.TestSecret, 2)
		s.Require().NoError(err)

		dbSecret, err := s.storage.GetSecret(id)

		s.Require().Equal(model.TestSecret.Info.TypeID, dbSecret.Info.TypeID)
		s.Require().Equal(model.TestSecret.Info.Title, dbSecret.Info.Title)
		s.Require().Equal(model.TestSecret.Info.Description, dbSecret.Info.Description)
		s.Require().Equal(model.TestSecret.Data, dbSecret.Data)
	})
}
