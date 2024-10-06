package repository_test

import (
	"context"
	"database/sql"
	"testing"
	"todo_list/internal/adapter/repository"
	"todo_list/internal/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestBasicUserOperations(t *testing.T) {
	ctx := context.Background()

	repo := repository.NewUsers()
	provider := cleanTablesAndCreateProvider(ctx, t)
	defer func() { _ = provider.Close() }()

	provider.ExecuteTx(ctx, func(ctx context.Context, connection domain.Connection) error {
		user := domain.User{
			ID:           domain.UserID(uuid.New()),
			Name:         "user name",
			Email:        "user@email.foo",
			PasswordHash: "some password hash",
			Token:        "some secret token",
		}
		require.NoError(t, repo.Create(ctx, connection, user))

		user.Name = "new user name"
		require.NoError(t, repo.Update(ctx, connection, user))

		newUser, err := repo.Read(ctx, connection, user.ID)
		require.NoError(t, err)
		require.Equal(t, user.Name, newUser.Name)

		require.NoError(t, repo.Delete(ctx, connection, user.ID))

		_, err = repo.Read(ctx, connection, user.ID)
		require.ErrorIs(t, err, sql.ErrNoRows)

		return nil
	})
}
