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
	user1, _ := userDomain.NewUser("lastName", "firstName", "user@example.com", "09000000000", "東京都", "渋谷区", "1-1-1")
	user2, _ := userDomain.NewUser("lastName2", "firstName2", "user2@example.com", "09000000001", "東京都", "新宿区", "1-1-1")
	transactionManager := NewTransactionManager()
	userRepository := NewUserRepository()
	ctx := context.Background()

	t.Run("正常系:トランザクションが完了し保存できること", func(t *testing.T) {
		var want1 *userDomain.User
		var want2 *userDomain.User
		err := transactionManager.RunInTransaction(ctx, func(ctx context.Context) error {
			userRepository.Save(ctx, user1)
			userRepository.Save(ctx, user2)
			return nil
		})
		want1, _ = userRepository.FindById(ctx, user1.ID())
		want2, _ = userRepository.FindById(ctx, user2.ID())

		if err != nil {
			t.Error(err)
		}
		if err != nil {
			t.Error(err)
		}
		if diff := cmp.Diff(want1.ID(), user1.ID()); diff != "" {
			t.Errorf("FindById() mismatch (-want +got):\n%s", diff)
		}
		if diff := cmp.Diff(want2.ID(), user2.ID()); diff != "" {
			t.Errorf("FindById() mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("異常系:ロールバックされること", func(t *testing.T) {
		var want1 *userDomain.User
		var want2 *userDomain.User
		err := transactionManager.RunInTransaction(ctx, func(ctx context.Context) error {
			userRepository.Save(ctx, user1)
			userRepository.Save(ctx, user2)
			err := ErrorRepositorySave(ctx, user1)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			t.Logf("トランザクション内Error: %v", err)
		}

		want1, _ = userRepository.FindById(ctx, user1.ID())
		if want1 != nil {
			t.Errorf("ロールバックされていない")
		}
		want2, _ = userRepository.FindById(ctx, user2.ID())
		if want2 != nil {
			t.Errorf("ロールバックされていない")
		}
	})
}

func ErrorRepositorySave(ctx context.Context, u *userDomain.User) error {
	return fmt.Errorf("明示的なエラー")
}
