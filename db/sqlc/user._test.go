package db

import (
	"context"
	"github.com/JCSong-89/go-udemy-master-class/utill"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomUser(t *testing.T) User {
	password, err := utill.HashPassword(utill.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       utill.RandomOwner(),
		HashedPassword: password,
		FullName:       utill.RandomOwner(),
		Email:          utill.RandomEmail(), // TODO: RandomEmail() is not implemented yet
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotNil(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChainedAt.IsZero())

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.CreatedAt, user2.CreatedAt)
	require.Equal(t, user1.PasswordChainedAt, user2.PasswordChainedAt)
}
