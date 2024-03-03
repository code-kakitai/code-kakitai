package api_test

import (
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db"
	dbTest "github/code-kakitai/code-kakitai/infrastructure/mysql/db/db_test"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db/dbgen"
	"github/code-kakitai/code-kakitai/presentation/settings"
	"github/code-kakitai/code-kakitai/server/route"
	"testing"

	"gopkg.in/testfixtures.v2"
)

var (
	fixtures *testfixtures.Context
)

func TestMain(m *testing.M) {
	var err error

	resource, pool := dbTest.CreateContainer()
	defer dbTest.CloseContainer(resource, pool)

	dbCon := dbTest.ConnectDB(resource, pool)
	defer dbCon.Close()

	dbTest.SetupTestDB("../../infrastructure/mysql/db/schema/schema.sql")
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

	api := settings.NewGinEngine()
	route.InitRoute(api)

	m.Run()
}
