package access_token

import (
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

func GetNewAccessToken() AccessToken {
	return AccessToken{
		ExpiresAt: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (token AccessToken) IsExpired() bool {
	expirationTime := time.Unix(token.ExpiresAt, 0)
	return expirationTime.Before(time.Now().UTC())
}