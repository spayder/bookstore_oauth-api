package access_token

import (
	"githab.com/spayder/bookstore_oauth-api/src/utils/errors"
	"strings"
	"time"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId int64 `json:"user_id"`
	ClientId int64 `json:"client_id"`
	ExpiresAt int64 `json:"expires_at"`
}

func (token *AccessToken) Validate() *errors.RestErr {
	token.AccessToken = strings.TrimSpace(token.AccessToken)
	if len(token.AccessToken) == 0 {
		return errors.BadRequestError("empty access token")
	}
	if token.UserId <= 0 {
		return errors.BadRequestError("invalid user id")
	}
	if token.ClientId <= 0 {
		return errors.BadRequestError("invalid client id")
	}
	if token.ExpiresAt <= 0 {
		return errors.BadRequestError("invalid expires at")
	}
	return nil
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		ExpiresAt: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (token AccessToken) IsExpired() bool {
	expirationTime := time.Unix(token.ExpiresAt, 0)
	return expirationTime.Before(time.Now().UTC())
}