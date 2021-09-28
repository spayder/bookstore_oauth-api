package db

import (
	"githab.com/spayder/bookstore_oauth-api/src/domain/access_token"
	"githab.com/spayder/bookstore_oauth-api/src/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}
type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {

}

func (r *dbRepository) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.InternalServerError("database connection not implemented yet")
}