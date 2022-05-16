package sqllite

import (
	"advdiploma/client/model"
	"advdiploma/client/storage"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var _ storage.Storage = (*Storage)(nil)

const secretsTbl string = `
CREATE TABLE IF NOT EXISTS secrets (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    userID INT,
    typeID INT,
	title TEXT NOT NULL,
	description TEXT,
	data TEXT NOT NULL
  );`

type Storage struct {
	db *sql.DB
}

//  NewStorage inits new connection to psql storage.
func NewStorage(file string) (*Storage, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	if _, err = db.Exec(secretsTbl); err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil
}

//  AddSecret adds new secret to storage
func (s *Storage) AddSecret(v model.Secret, userID int) (int64, error) {
	stmt, err := s.db.Prepare("INSERT INTO secrets(userID, typeID, title, description, data) VALUES(?,?,?,?,?)")
	if err != nil {
		return 0, err
	}

	r, err := stmt.Exec(userID, v.Info.TypeID, v.Info.Title, v.Info.Description, v.Data)
	if err != nil {
		return 0, err
	}

	id, err := r.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

//  GetSecret returns secret from storage
func (s *Storage) GetSecret(id int64) (model.Secret, error) {
	res := model.Secret{Info: model.Info{}}
	if err := s.db.QueryRow(
		"SELECT id, typeID, title, description, data FROM secrets WHERE id=@id",
		sql.Named("id", id),
	).Scan(
		&res.Info.ID,
		&res.Info.TypeID,
		&res.Info.Title,
		&res.Info.Description,
		&res.Data); err != nil {
		return model.Secret{}, err
	}
	return res, nil
}

//  GetInfoForUser returns array of info for user secrets
func (s *Storage) GetInfoForUser(userID int) ([]model.Info, error) {
	var infos []model.Info

	rows, err := s.db.Query(
		"SELECT id, typeID, title, description FROM secrets WHERE userID=@userID",
		sql.Named("userID", userID))

	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println(err.Error())
		}
	}()

	for rows.Next() {
		var el model.Info
		err = rows.Scan(&el.ID, &el.TypeID, &el.Title, &el.Description)
		if err != nil {
			return nil, err
		}

		infos = append(infos, el)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	if len(infos) == 0 {
		infos = make([]model.Info, 0)
	}

	return infos, nil
}

//  Close  closes database connection.
func (s *Storage) Close() {
	if s.db == nil {
		return
	}

	if err := s.db.Close(); err != nil {
		log.Println(err.Error())
	}
	s.db = nil
}
