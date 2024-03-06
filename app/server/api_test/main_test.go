package api_test

import (
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db"
	dbTest "github/code-kakitai/code-kakitai/infrastructure/mysql/db/db_test"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db/dbgen"
	"github/code-kakitai/code-kakitai/presentation/settings"
	"github/code-kakitai/code-kakitai/server/route"
	"testing"

	"github.com/gin-gonic/gin"
	"gopkg.in/testfixtures.v2"
)

var (
	fixtures *testfixtures.Context
	api      *gin.Engine
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
	dbTest.SetupTestDB("../../infrastructure/mysql/db/schema/schema.sql")

	// テストデータの準備
	fixtures, err = testfixtures.NewFolder(
		dbCon,
		&testfixtures.MySQL{},
		"../../infrastructure/mysql/fixtures",
	)
	if err != nil {
		panic(err)
	}

	q := dbgen.New(dbCon)
	db.SetReadQuery(q)
	db.SetDB(dbCon)

	api = settings.NewGinEngine()
	route.InitRoute(api)

	m.Run()
}

func resetTestData(t *testing.T) {
	t.Helper()
	if err := fixtures.Load(); err != nil {
		t.Fatal(err)
	}
}
