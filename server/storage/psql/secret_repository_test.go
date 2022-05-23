package psql

import (
	"advdiploma/server/model"
	"github.com/google/uuid"
	"github.com/icrowley/fake"
	"time"
)

func (s *TestSuite) TestSecret_AddGet() {

	s.Run("Create non-existing user", func() {
		secret := model.Secret{
			UserID:   uuid.New(),
			DeviceID: uuid.New(),
			Data:     fake.CharactersN(6000),
		}

		_, err := s.storage.Secret().Add(s.ctx, secret)
		s.Require().Error(err)
	})

	s.Run("Create and Get", func() {
		user, err := s.storage.User().Create(s.ctx, model.User{
			Login:        fake.CharactersN(60),
			PasswordHash: fake.CharactersN(60),
		})
		s.Require().NoError(err)

		secret := model.Secret{
			UserID:     user.ID,
			DeviceID:   uuid.New(),
			Data:       fake.CharactersN(6000),
			IsDeleted:  false,
			UploadedAt: time.Now(),
			DeletedAt:  time.Now(),
		}

		secretAdded, err := s.storage.Secret().Add(s.ctx, secret)
		s.Require().NoError(err)

		secretRes, err := s.storage.Secret().Get(s.ctx, secretAdded.ID)
		s.Require().NoError(err)

		s.Assert().EqualValues(secretAdded.ID, secretRes.ID)

		s.Assert().EqualValues(secret.UserID, secretRes.UserID)
		s.Assert().EqualValues(secret.DeviceID, secretRes.DeviceID)
		s.Assert().EqualValues(secret.Data, secretRes.Data)
		s.Assert().EqualValues(secret.IsDeleted, secretRes.IsDeleted)

	})

}
