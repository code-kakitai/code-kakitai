package repository

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/go-cmp/cmp"

	userDomain "github/code-kakitai/code-kakitai/domain/user"
)

func TestUserRepository_FindById(t *testing.T) {
	user, err := userDomain.Reconstruct("01HCNYK0PKYZWB0ZT1KR0EPWGP", "example@test.com", "08011112222", "山田", "太郎", "東京都", "渋谷区", "1-1-1")
	if err != nil {
		t.Error(err)
	}
	tests := []struct {
		name string
		want *userDomain.User
	}{
		{
			name: "IDによってユーザーが取得ができること",
			want: user,
		},
	}
	userRepository := NewUserRepository()
	ctx := context.Background()
	resetTestData(t)
	for _, tt := range tests {
		t.Run(fmt.Sprintf(": %s", tt.name), func(t *testing.T) {
			got, err := userRepository.FindById(ctx, "01HCNYK0PKYZWB0ZT1KR0EPWGP")
			if err != nil {
				t.Error(err)
			}
			if diff := cmp.Diff(got.ID(), tt.want.ID()); diff != "" {
				t.Errorf("FindById() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestUserRepository_Save(t *testing.T) {
	user, _ := userDomain.NewUser("test@example.com", "09000000000", "lastName", "firstName", "東京都", "渋谷区", "1-1-1")
	tests := []struct {
		name  string
		input *userDomain.User
		want  *userDomain.User
	}{
		{
			name:  "保存かつ取得ができること",
			input: user,
			want:  user,
		},
	}
	userRepository := NewUserRepository()
	ctx := context.Background()
	resetTestData(t)
	for _, tt := range tests {
		t.Run(fmt.Sprintf(": %s", tt.name), func(t *testing.T) {
			err := userRepository.Save(ctx, tt.input)
			if err != nil {
				t.Error(err)
			}

			got, err := userRepository.FindById(ctx, tt.input.ID())
			if err != nil {
				t.Error(err)
			}

			if diff := cmp.Diff(got.ID(), tt.want.ID()); diff != "" {
				t.Errorf("FindById() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
