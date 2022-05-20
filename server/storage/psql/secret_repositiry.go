package psql

import (
	"advdiploma/server/model"
	"context"
	"database/sql"
)

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
	stmt, err := s.db.Prepare(
		"INSERT INTO secrets(user_id, client_id, data) VALUES(@user_id,@client_id,@data) RETURNING id, user_id, client_id, 'data'")
	if err != nil {
		return model.Secret{}, err
	}

	if err := stmt.QueryRowContext(
		ctx,
		sql.Named("user_id", secret.UserID),
		sql.Named("client_id", secret.ClientID),
		sql.Named("data", secret.Data),
	).Scan(
		&secret.ID,
		&secret.UserID,
		&secret.ClientID,
		&secret.Data); err != nil {

		return model.Secret{}, err
	}

	return secret, nil
}

func (s *secretRepository) Get(ctx context.Context, id int64) (model.Secret, error) {
	res := model.Secret{}
	if err := s.db.QueryRowContext(ctx,
		"SELECT id, user_id, client_id, data FROM secrets WHERE id=@id",
		sql.Named("id", id),
	).Scan(
		&res.ID,
		&res.UserID,
		&res.ClientID,
		&res.Data); err != nil {
		return model.Secret{}, err
	}

	return res, nil
}
