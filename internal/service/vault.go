package service

import (
	"github.com/esceer/vault/internal/adapter"
	apimodel "github.com/esceer/vault/internal/api/model"
	"github.com/esceer/vault/internal/common"
	"github.com/esceer/vault/internal/security"
	"github.com/esceer/vault/internal/storage"
	dbmodel "github.com/esceer/vault/internal/storage/model"
)

type VaultService interface {
	GetAll() ([]*apimodel.CredentialResponse, error)
	GetSecret(common.Identifier, common.Secret) (common.Secret, error)
	Save(*apimodel.CredentialCreate) error
	Delete(common.Identifier) error
}

type vaultService struct {
	store storage.CredentialStore
}

func NewVaultService(s storage.CredentialStore) VaultService {
	return &vaultService{s}
}

func (s vaultService) GetAll() ([]*apimodel.CredentialResponse, error) {
	creds, err := s.store.GetAll()
	return adapter.DbSliceToApiSlice(creds), err
}

func (s vaultService) GetSecret(id common.Identifier, masterKey common.Secret) (common.Secret, error) {
	cred, err := s.store.GetById(id)
	if err != nil {
		return nil, err
	}

	decoded, err := s.decodeSecret(masterKey, cred)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

func (s vaultService) Save(apiCred *apimodel.CredentialCreate) error {
	encoded, err := s.encodeSecret(apiCred)
	if err != nil {
		return err
	}

	dbCred := &dbmodel.Credential{
		User:   apiCred.User,
		Site:   apiCred.Site,
		Secret: encoded,
	}
	return s.store.Save(dbCred)
}

func (vaultService) encodeSecret(cred *apimodel.CredentialCreate) (common.Secret, error) {
	hash := security.Hash32(security.Hash32(cred.MasterKey, []byte(cred.User)), []byte(cred.Site))
	return security.Encode(hash, cred.Secret)
}

func (vaultService) decodeSecret(masterKey common.Secret, cred *dbmodel.Credential) (common.Secret, error) {
	hash := security.Hash32(security.Hash32(masterKey, []byte(cred.User)), []byte(cred.Site))
	return security.Decode(hash, cred.Secret)
}

func (s vaultService) Delete(id common.Identifier) error {
	return s.store.Delete(id)
}
