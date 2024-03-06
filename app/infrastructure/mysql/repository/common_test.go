package repository

import (
	"testing"

	"gopkg.in/testfixtures.v2"

	"github/code-kakitai/code-kakitai/infrastructure/mysql/db"
	dbTest "github/code-kakitai/code-kakitai/infrastructure/mysql/db/db_test"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db/dbgen"
)

var (
	fixtures *testfixtures.Context
)

func TestMain(m *testing.M) {
	var err error

	// DBの立ち上げ
	resource, pool := dbTest.CreateContainer()
	defer dbTest.CloseContainer(resource, pool)

	// DBへ接続する
	dbCon := dbTest.ConnectDB(resource, pool)
	defer dbCon.Close()

	// テスト用DBをセットアップ
	dbTest.SetupTestDB("../db/schema/schema.sql")

	// テストデータの準備
	fixturePath := "../fixtures"
	fixtures, err = testfixtures.NewFolder(dbCon, &testfixtures.MySQL{}, fixturePath)
	if err != nil {
		panic(err)
	}

	q := dbgen.New(dbCon)
	db.SetQuery(q)
	db.SetDB(dbCon)

	// テスト実行
	m.Run()
}

func resetTestData(t *testing.T) {
	t.Helper()
	if err := fixtures.Load(); err != nil {
		t.Fatal(err)
	}
}
