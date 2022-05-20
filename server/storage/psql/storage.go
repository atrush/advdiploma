package psql

import (
	"advdiploma/server/storage"
	"database/sql"
	"fmt"
	"log"
)

var _ storage.Storage = (*Storage)(nil)

type Storage struct {
	secretRepo   *secretRepository
	userRepo     *userRepository
	db           *sql.DB
	conStringDSN string
}

//  NewStorage inits new connection to psql storage.
//  !!!! On init drop all and init tables.
func NewStorage(dsn string) (*Storage, error) {
	if dsn == "" {
		return nil, fmt.Errorf("error init data base:%v", "dsn string is empty")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	//if err := initBase(db); err != nil {
	//	return nil, err
	//}

	st := &Storage{
		db:           db,
		conStringDSN: dsn,
	}

	st.secretRepo = newSecretRepository(db)
	st.userRepo = newUserRepository(db)

	return st, nil
}

//  User returns users repository.
func (s *Storage) User() storage.UserRepository {
	return s.userRepo
}

//  Secret returns users repository.
func (s *Storage) Secret() storage.SecretRepository {
	return s.secretRepo
}

//  Close  closes database connection.
func (s Storage) Close() {
	if s.db == nil {
		return
	}

	if err := s.db.Close(); err != nil {
		log.Println(err.Error())
	}

	s.db = nil
}
