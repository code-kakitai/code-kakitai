package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db/dbgen"
	"log"
	"strconv"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/k0kubun/sqldef"
	"github.com/k0kubun/sqldef/database"
	"github.com/k0kubun/sqldef/database/mysql"
	"github.com/k0kubun/sqldef/parser"
	"github.com/k0kubun/sqldef/schema"
	"github.com/ory/dockertest"
)

type DBContainer struct {
	db       *sql.DB
	query    *dbgen.Queries
	resource *dockertest.Resource
	pool     *dockertest.Pool
	port     int
}

var dbContainer *DBContainer

func (dc *DBContainer) Close() {
	fmt.Println("container close")
	if err := dc.pool.Purge(dc.resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

}

func NewDBContainer() error {
	dbContainer = &DBContainer{}
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	dbContainer.resource, dbContainer.pool = createContainer()
	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := dbContainer.pool.Retry(func() error {
		var err error
		dbContainer.port, err = strconv.Atoi(dbContainer.resource.GetPort("3306/tcp"))
		if err != nil {
			log.Fatalf("Could not parse port: %s", err)
			return err
		}
		dbContainer.db, err = sql.Open("mysql", fmt.Sprintf("root:secret@(localhost:%s)/code-kakitai", dbContainer.resource.GetPort("3306/tcp")))
		if err != nil {
			log.Fatalf("SQL Open Error: %s", err)
			return err
		}
		return dbContainer.db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
		return err
	}
	dbContainer.query = dbgen.New(dbContainer.db)
	return nil
}

func createContainer() (*dockertest.Resource, *dockertest.Pool) {
	pool, err := dockertest.NewPool("")
	pool.MaxWait = time.Minute * 2
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// Dockerコンテナ起動時の細かいオプションを指定する
	// テーブル定義などはここで流し込むのが良さそう?
	runOptions := &dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "8.0",
		Env:        []string{"MYSQL_ROOT_PASSWORD=secret", "MYSQL_DATABASE=code-kakitai"},
		Mounts:     []string{},
		Cmd: []string{
			"mysqld",
			"--character-set-server=utf8mb4",
			"--collation-server=utf8mb4_unicode_ci",
		},
	}

	// コンテナを起動
	resource, err := pool.RunWithOptions(runOptions)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	return resource, pool
}

func TestMain(m *testing.M) {
	err := NewDBContainer()
	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}
	defer dbContainer.Close()

	// マイグレーション
	desiredDDLs, err := sqldef.ReadFile("../db/schema/schema.sql")
	if err != nil {
		log.Fatalf("failed to read schema file: %s", err)
	}
	options := &sqldef.Options{DesiredDDLs: desiredDDLs}
	sp := database.NewParser(parser.ParserModeMysql)
	database, err := mysql.NewDatabase(database.Config{
		Host:     "127.0.0.1",
		Port:     dbContainer.port,
		User:     "root",
		Password: "secret",
		DbName:   "code-kakitai",
	})
	if err != nil {
		log.Fatal(err)
	}
	// end マイグレーション
	sqldef.Run(schema.GeneratorModeMysql, database, sp, options)
	m.Run()
}

func TestSomething(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "テスト",
			input:    "テスト",
			expected: "テスト",
		},
	}
	for _, td := range tests {
		userRepository := NewPlayerRepositoryImpl(dbContainer.query)
		user, err := userRepository.FindById(context.Background(), "1")
		if err != nil {
			t.Error(err)
		}

		t.Run(fmt.Sprintf(": %s", td.name), func(t *testing.T) {
			if user == nil {
				t.Error("ユーザは存在していない")
			} else {
				t.Log("ユーザーは存在している")
			}
		})
	}
}
