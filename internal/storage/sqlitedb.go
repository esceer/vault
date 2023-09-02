package storage

import (
	"database/sql"
	"errors"
)

type dbStorage struct {
	db *sql.DB
}

func NewDBStorage(db *sql.DB) Store {
	return &dbStorage{db: db}
}

func (s *dbStorage) GetKeysByUser(user string) ([]string, error) {
	rows, err := s.db.Query("select key from safe where user=?", user)
	if err != nil {
		return nil, err
	}

	var keys []string
	for rows.Next() {
		var key string
		rows.Scan(&key)
		keys = append(keys, key)
	}

	return keys, nil
}

func (s *dbStorage) Store(user, key string, secret []byte) error {
	stmt, err := s.db.Prepare("insert into safe(key, secret) values(?,?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(key, secret)
	return err
}

func (s *dbStorage) GetSecret(id identifier) ([]byte, error) {
	var secret []byte

	if err := s.db.QueryRow("select secret from safe where id=?", id).Scan(&secret); err != nil {
		return nil, err
	}
	return secret, nil
}

func (s *dbStorage) Delete(id identifier) error {
	stmt, err := s.db.Prepare("delete from safe where key=?")
	if err != nil {
		return err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected < 1 {
		return errors.New("entry not found")
	}
	if affected > 1 {
		return errors.New("multiple entries found")
	}

	return nil
}
