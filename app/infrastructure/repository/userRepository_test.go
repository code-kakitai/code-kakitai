package repository

import (
	"context"
	"fmt"
	userDomain "github/code-kakitai/code-kakitai/domain/user"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	ud1, _ := userDomain.BuilderUser(
		"1",
		"test@example.com",
		"09012345678",
		"test",
		"test",
		"東京都",
		"渋谷区",
		"1-1-1",
	)

	tests := []struct {
		name  string
		id    string
		input *userDomain.User
	}{
		{
			name:  "プレイヤーが作成されること",
			id:    "1",
			input: ud1,
		},
	}
	for _, td := range tests {
		t.Run(fmt.Sprintf(": %s", td.name), func(t *testing.T) {
			ctx := context.Background()
			repository := NewUserRepositoryImpl(Queries)
			us, err := repository.FindById(ctx, td.id)
			fmt.Printf("user: %v\n", us)
			if err != nil {
				t.Fatal(err)
			}

			// assert
			// assert.Equal(t, result, td.input)
		})
	}
}
