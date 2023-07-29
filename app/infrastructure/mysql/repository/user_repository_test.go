package repository

import (
	"context"
	"fmt"
	userDomain "github/code-kakitai/code-kakitai/domain/user"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/go-cmp/cmp"
)

func TestSomething(t *testing.T) {
	user, _ := userDomain.NewUser("lastName", "firstName", "tomoya.kamaji@gmail.com", "09071121428", "東京都", "渋谷区", "1-1-1")
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
	for _, tt := range tests {
		userRepository := NewUserRepository(query)
		ctx := context.Background()
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
