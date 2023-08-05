package repository

import (
	"context"
	"fmt"
	"github/code-kakitai/code-kakitai/application/transaction"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db/dbgen"
	"github/code-kakitai/code-kakitai/util"
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
	db.SetQuery(q)
	defer db.SetQuery(dbgen.New(dbcon)) // トランザクション終了後にQueriesを元に戻す

	// QueriesをContextにセット
	ctxWithQueries := util.WithQueries(ctx, q)

	// トランザクション内の関数を実行
	err = fn(ctxWithQueries)
	if err != nil {
		// トランザクション内でエラーが発生したらロールバック
		fmt.Printf("db rollback: %v\n", err)
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
