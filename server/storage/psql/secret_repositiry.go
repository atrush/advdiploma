package psql

import (
	"advdiploma/server/model"
	"advdiploma/server/storage"
	"context"
	"database/sql"
	"github.com/google/uuid"
)

var _ storage.SecretRepository = (*secretRepository)(nil)

//  secretRepository implements SecretRepository interface, provides actions with order records in psql storage.
type secretRepository struct {
	db *sql.DB
}

//  newOrderRepository inits new order repository.
func newSecretRepository(db *sql.DB) *secretRepository {
	return &secretRepository{
		db: db,
	}
}

func (s *secretRepository) Add(ctx context.Context, secret model.Secret) (model.Secret, error) {
	err := s.db.QueryRowContext(
		ctx,
		"INSERT INTO secrets(user_id,device_id,data) VALUES($1,$2,$3) "+
			"RETURNING id, user_id, device_id, is_deleted ,data, created_at,deleted_at",
		secret.UserID,
		secret.DeviceID,
		secret.Data,
	).Scan(
		&secret.ID,
		&secret.UserID,
		&secret.DeviceID,
		&secret.IsDeleted,
		&secret.Data,
		&secret.UploadedAt,
		&secret.DeletedAt,
	)

	if err != nil {
		////  if exist return ErrorConflictSaveUser
		//pqErr, ok := err.(*pq.Error)
		//if ok && pqErr.Code == pgerrcode.UniqueViolation && pqErr.Constraint == "users_login_key" {
		//	return model.User{}, model.ErrorConflictSaveUser
		//}
		return model.Secret{}, err
	}

	return secret, nil
}

func (s *secretRepository) Get(ctx context.Context, id uuid.UUID) (model.Secret, error) {
	res := model.Secret{}
	if err := s.db.QueryRowContext(ctx,
		"SELECT id, user_id,device_id,is_deleted,data, created_at,deleted_at FROM secrets WHERE id=$1",
		id,
	).Scan(
		&res.ID,
		&res.UserID,
		&res.DeviceID,
		&res.IsDeleted,
		&res.Data,
		&res.UploadedAt,
		&res.DeletedAt,
	); err != nil {
		return model.Secret{}, err
	}

	return res, nil
}
