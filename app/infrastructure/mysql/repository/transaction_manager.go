package repository

import (
	"context"
	"fmt"
	"log"

	"github/code-kakitai/code-kakitai/application/transaction"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db/dbgen"
)

type TransactionManager struct{}

func NewTransactionManager() transaction.TransactionManager {
	return &TransactionManager{}
}

func (tm *TransactionManager) RunInTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	// トランザクションを開始
	dbcon := db.GetDB()
	tx, err := dbcon.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	// トランザクション用のQueriesを作成
	q := dbgen.New(tx)

	// QueriesをContextにセット
	ctxWithQueries := db.WithQueries(ctx, q)

	// トランザクション内の関数を実行
	err = fn(ctxWithQueries)
	if err != nil {
		// トランザクション内でエラーが発生したらロールバック
		log.Printf("db rollback: %v\n", err)
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	// 全ての関数が成功したらコミット
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit failed: %w", err)
	}

	return nil
}
