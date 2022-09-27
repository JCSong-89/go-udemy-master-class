package utill

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)
	hash, err := HashPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, hash)

	if err != nil {
		t.Error(err)
	}

	result, err := CheckPasswordHash(password, hash)
	require.True(t, result)
	require.NoError(t, err)

	wrongPassword := RandomString(6)
	result, err = CheckPasswordHash(wrongPassword, hash)
	require.False(t, result)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
