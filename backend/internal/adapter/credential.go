package adapter

import (
	apimodel "github.com/esceer/vault/backend/internal/api/model"
	dbmodel "github.com/esceer/vault/backend/internal/storage/model"
)

func DbToApi(dbModel *dbmodel.Credential) *apimodel.CredentialResponse {
	if dbModel == nil {
		return nil
	}
	return &apimodel.CredentialResponse{
		ID:        dbModel.ID,
		User:      dbModel.User,
		Site:      dbModel.Site,
		CreatedAt: dbModel.CreatedAt,
	}
}

func DbSliceToApiSlice(dbModels []*dbmodel.Credential) []*apimodel.CredentialResponse {
	apiModel := make([]*apimodel.CredentialResponse, len(dbModels))
	for i, dbm := range dbModels {
		apiModel[i] = DbToApi(dbm)
	}
	return apiModel
}
