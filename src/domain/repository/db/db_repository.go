package db

import (
	"githab.com/spayder/bookstore_oauth-api/src/clients/cassandra"
	"githab.com/spayder/bookstore_oauth-api/src/domain/access_token"
	"githab.com/spayder/bookstore_oauth-api/src/utils/errors"
	"github.com/gocql/gocql"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}
type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(token access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(token access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {

}

func (r *dbRepository) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}
	defer session.Close()

	var result access_token.AccessToken
	query := "SELECT access_token, user_id, client_id, expires_at FROM access_tokens WHERE access_token = ?;"
	if err := session.Query(query, accessTokenId).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.ExpiresAt,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NotFoundError("no access token found with given id")
		}
		return nil, errors.InternalServerError(err.Error())
	}

	return &result, nil
}

func (r *dbRepository) Create(token access_token.AccessToken)  *errors.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer session.Close()

	query := "INSERT INTO access_tokens (access_token, user_id, client_id, expires_at) " +
		"VALUES (?, ?, ?, ?);"

	if err := session.Query(query, token.AccessToken, token.UserId, token.ClientId, token.ExpiresAt).Exec(); err != nil {
		return errors.InternalServerError(err.Error())
	}

	return nil
}

func (r *dbRepository) UpdateExpirationTime(token access_token.AccessToken) *errors.RestErr {
	session, err := cassandra.GetSession()
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer session.Close()

	query := "UPDATE access_tokens SET expires_at = ? WHERE access_token = ?;"

	if err := session.Query(query, token.ExpiresAt, token.AccessToken).Exec(); err != nil {
		return errors.InternalServerError(err.Error())
	}

	return nil
}