package storage

import (
	"github.com/esceer/vault/internal/common"
	"github.com/esceer/vault/internal/storage/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CredentialStore interface {
	GetAll() ([]*model.Credential, error)
	GetById(common.Identifier) (*model.Credential, error)
	Save(*model.Credential) error
	Delete(common.Identifier) error
}

type dbStore struct {
	db *gorm.DB
}

func NewDBStore(db *gorm.DB) CredentialStore {
	return &dbStore{db: db}
}

func (s *dbStore) GetAll() ([]*model.Credential, error) {
	var creds []*model.Credential
	result := s.db.Omit("secret").Find(&creds)
	return creds, result.Error
}

func (s *dbStore) GetById(id common.Identifier) (*model.Credential, error) {
	var cred model.Credential
	result := s.db.Preload(clause.Associations).Omit("password").First(&cred, id)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, ErrNotFound
	}
	return &cred, result.Error
}

func (s *dbStore) Save(cred *model.Credential) error {
	return s.db.Create(cred).Error
}

func (s *dbStore) Delete(id common.Identifier) error {
	return s.db.Delete(&model.Credential{}, id).Error
}
