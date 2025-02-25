package repository

import (
	"context"
	"hacktiv/model"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestUserRepository(t *testing.T) {
	db, mock := NewMockDB()

	userRepo := NewUserRepository(db)

	boolTrue := true

	user := model.User{
		ID:         1,
		Name:       "John Doe",
		Email:      "alif@go.dev",
		ActiveUser: &boolTrue,
		CreatedAt:  time.Now(),
	}

	t.Run("CreateUser", func(t *testing.T) {
		mock.ExpectBegin().WillReturnError(nil)
		mock.ExpectQuery("^INSERT INTO \"users\" (.*)").
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit().WillReturnError(nil)

		err := userRepo.CreateUser(context.Background(), user)
		require.NoError(t, err)
	})

	t.Run("GetAllUsers", func(t *testing.T) {
		mock.ExpectQuery(`^SELECT (.*) FROM "users"`).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email"}).
				AddRow(1, "John Doe", "john@example.com"))

		users, err := userRepo.GetAllUsers(context.Background())
		require.NoError(t, err)
		require.Equal(t, 1, len(users))
		require.Equal(t, user.ID, users[0].ID)
	})
}
