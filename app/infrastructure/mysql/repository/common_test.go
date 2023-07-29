package repository

import (
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db/db_test"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db/dbgen"
	"testing"
)

var query *dbgen.Queries

func TestMain(m *testing.M) {
	// DBの立ち上げ
	resource, pool := db_test.CreateContainer()
	defer db_test.CloseContainer(resource, pool)

	// DBへ接続する
	db := db_test.ConnectDB(resource, pool)
	defer db.Close()

	// テスト用DBをセットアップ
	db_test.SetupTestDB()

	query = dbgen.New(db)

	// テスト実行
	m.Run()
}
