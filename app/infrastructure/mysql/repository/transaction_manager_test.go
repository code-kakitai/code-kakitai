package repository

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/go-cmp/cmp"

	userDomain "github/code-kakitai/code-kakitai/domain/user"
)

func TestTransactionManager(t *testing.T) {
	transactionManager := NewTransactionManager()
	userRepository := NewUserRepository()
	ctx := context.Background()

	t.Run("正常系:トランザクションが完了し保存できること", func(t *testing.T) {
		user1, _ := userDomain.NewUser("user@example.com", "09000000000", "lastName", "firstName", "東京都", "渋谷区", "1-1-1")
		user2, _ := userDomain.NewUser("user2@example.com", "09000000001", "lastName2", "firstName2", "東京都", "新宿区", "1-1-1")
		transactionManager.RunInTransaction(ctx, func(ctx context.Context) error {
			userRepository.Save(ctx, user1)
			userRepository.Save(ctx, user2)
			return nil
		})
		want1, _ := userRepository.FindById(ctx, user1.ID())
		want2, _ := userRepository.FindById(ctx, user2.ID())
		if diff := cmp.Diff(want1.ID(), user1.ID()); diff != "" {
			t.Errorf("FindById() mismatch (-want +got):\n%s", diff)
		}
		if diff := cmp.Diff(want2.ID(), user2.ID()); diff != "" {
			t.Errorf("FindById() mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("異常系:トランザクション内でエラーが発生した際はロールバックされること", func(t *testing.T) {
		user1, _ := userDomain.NewUser("user@example.com", "09000000000", "lastName", "firstName", "東京都", "渋谷区", "1-1-1")
		user2, _ := userDomain.NewUser("user2@example.com", "09000000001", "lastName2", "firstName2", "東京都", "新宿区", "1-1-1")
		transactionManager.RunInTransaction(ctx, func(ctx context.Context) error {
			userRepository.Save(ctx, user1)
			userRepository.Save(ctx, user2)
			err := errorRepositorySave(ctx, user1)
			if err != nil {
				return err
			}
			return nil
		})

		want1, _ := userRepository.FindById(ctx, user1.ID())
		if want1 != nil {
			t.Errorf("user1が保存されている(ロールバックされていない)")
		}
		want2, _ := userRepository.FindById(ctx, user2.ID())
		if want2 != nil {
			t.Errorf("user2が保存されている(ロールバックされていない)")
		}
	})
}

func errorRepositorySave(ctx context.Context, u *userDomain.User) error {
	return fmt.Errorf("明示的なエラー")
}
