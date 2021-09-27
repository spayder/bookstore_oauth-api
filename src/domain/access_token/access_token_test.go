package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestExpirationTimeConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "expiration time for token should be 24 hours")
}

func TestGetAccessToken(t *testing.T) {
	token := GetNewAccessToken()

	assert.False(t, token.IsExpired(), "token expiration date should not be in the past")
	assert.Emptyf(t, token.AccessToken, "token should not has defined access token id")
	assert.True(t, token.UserId == 0, "token should not has an associated user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	token := AccessToken{}

	assert.True(t, token.IsExpired(), "empty token should be expired by default")

	token.ExpiresAt = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, token.IsExpired(), "token expired three hours from now should not be expired")
}