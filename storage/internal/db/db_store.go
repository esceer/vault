package db

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

const (
	database = "vault.db"
	driver   = "sqlite3"
)

type dbStore struct {
}

func NewStore() *dbStore {
	return &dbStore{}
}

func (s *dbStore) ListKeys() ([]string, error) {
	db, err := sql.Open(driver, database)
	defer db.Close()

	if err != nil {
		return nil, err
	}

	rows, err := db.Query("select key from safe")
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

func (s *dbStore) Store(key string, secret []byte) error {
	db, err := sql.Open(driver, database)
	defer db.Close()

	if err != nil {
		return err
	}

	stmt, err := db.Prepare("insert into safe(key, secret) values(?,?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(key, secret)
	if err != nil {
		return err
	}

	return nil
}

func (s *dbStore) Retrieve(key string) ([]byte, error) {
	db, err := sql.Open(driver, database)
	defer db.Close()

	if err != nil {
		return nil, err
	}

	var secret []byte

	err = db.QueryRow("select secret from safe where key=?", key).Scan(&secret)
	if err != nil {
		return nil, err
	}

	return secret, nil
}

func (s *dbStore) Delete(key string) error {
	db, err := sql.Open(driver, database)
	defer db.Close()

	if err != nil {
		return err
	}

	stmt, err := db.Prepare("delete from safe where key=?")
	if err != nil {
		return err
	}

	result, err := stmt.Exec(key)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected < 1 {
		return errors.New("Entry not found")
	}
	if affected > 1 {
		return errors.New("Multiple entries found")
	}

	return nil
}
