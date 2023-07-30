package repository

import (
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db"
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
	dbCon := db_test.ConnectDB(resource, pool)
	defer dbCon.Close()

	// テスト用DBをセットアップ
	db_test.SetupTestDB()

	q := dbgen.New(dbCon)
	db.SetQuery(q)

	// テスト実行
	m.Run()
}
