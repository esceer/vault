package storage

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type dbStore struct {
	db *gorm.DB
}

func NewDBStore(db *gorm.DB) Store {
	return &dbStore{db: db}
}

func (s *dbStore) GetKeysByUser(user string) ([]string, error) {
	var keys []string
	result := s.db.Model(&Credential{}).Where("user=?", user).Pluck("key", &keys)
	return keys, result.Error
}

func (s *dbStore) GetById(id identifier) (*Credential, error) {
	var cred Credential
	result := s.db.Preload(clause.Associations).Omit("Password").First(&cred, id)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, ErrNotFound
	}
	return &cred, result.Error
}

func (s *dbStore) Save(cred *Credential) error {
	return s.db.Create(cred).Error
}

func (s *dbStore) Delete(id identifier) error {
	return s.db.Delete(&Credential{}, id).Error
}
