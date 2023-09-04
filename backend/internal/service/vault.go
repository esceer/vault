package service

import (
	"github.com/esceer/vault/backend/internal/adapter"
	apimodel "github.com/esceer/vault/backend/internal/api/model"
	"github.com/esceer/vault/backend/internal/common"
	"github.com/esceer/vault/backend/internal/security"
	"github.com/esceer/vault/backend/internal/storage"
	dbmodel "github.com/esceer/vault/backend/internal/storage/model"
)

type VaultService interface {
	GetAll() ([]*apimodel.CredentialResponse, error)
	GetSecret(common.Identifier, string) (*apimodel.SecretResponse, error)
	Save(*apimodel.CredentialCreate, string) error
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

func (s vaultService) GetSecret(id common.Identifier, masterkey string) (*apimodel.SecretResponse, error) {
	cred, err := s.store.GetById(id)
	if err != nil {
		return nil, err
	}

	decoded, err := s.decodeSecret(cred, masterkey)
	if err != nil {
		return nil, err
	}
	return &apimodel.SecretResponse{Secret: decoded}, nil
}

func (s vaultService) Save(apiCred *apimodel.CredentialCreate, masterkey string) error {
	encoded, err := s.encodeSecret(apiCred, masterkey)
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

func (vaultService) encodeSecret(cred *apimodel.CredentialCreate, masterkey string) ([]byte, error) {
	hash := security.Hash32(security.Hash32([]byte(masterkey), []byte(cred.User)), []byte(cred.Site))
	return security.Encode(hash, []byte(cred.Secret))
}

func (vaultService) decodeSecret(cred *dbmodel.Credential, masterkey string) (string, error) {
	hash := security.Hash32(security.Hash32([]byte(masterkey), []byte(cred.User)), []byte(cred.Site))
	if b, err := security.Decode(hash, cred.Secret); err != nil {
		return "", err
	} else {
		return string(b), nil
	}
}

func (s vaultService) Delete(id common.Identifier) error {
	return s.store.Delete(id)
}
