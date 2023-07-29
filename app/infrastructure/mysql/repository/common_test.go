package repository

import (
	"database/sql"
	"fmt"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db/dbgen"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/k0kubun/sqldef"
	"github.com/k0kubun/sqldef/database"
	"github.com/k0kubun/sqldef/database/mysql"
	"github.com/k0kubun/sqldef/parser"
	"github.com/k0kubun/sqldef/schema"
	"github.com/ory/dockertest"
)

var (
	username = "root"
	password = "secret"
	hostname = "localhost"
	dbName   = "code-kakitai"
	port     = 0 // コンテナ起動時に決定する
)

var query *dbgen.Queries

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
		Env: []string{
			"MYSQL_ROOT_PASSWORD=" + password,
			"MYSQL_DATABASE=" + dbName,
		},
		Mounts: []string{},
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

func closeContainer(resource *dockertest.Resource, pool *dockertest.Pool) {
	// コンテナの終了
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func connectDB(resource *dockertest.Resource, pool *dockertest.Pool) *sql.DB {
	// DB(コンテナ)との接続
	var db *sql.DB
	if err := pool.Retry(func() error {
		time.Sleep(time.Second * 3)
		var err error
		port, err = strconv.Atoi(resource.GetPort("3306/tcp"))
		if err != nil {
			return err
		}
		db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", username, password, hostname, resource.GetPort("3306/tcp"), dbName))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}
	return db
}

func setupTestDB() {
	// マイグレーション
	desiredDDLs, err := sqldef.ReadFile("../db/schema/schema.sql")
	if err != nil {
		log.Fatalf("failed to read schema file: %s", err)
	}
	options := &sqldef.Options{DesiredDDLs: desiredDDLs}
	sp := database.NewParser(parser.ParserModeMysql)
	database, err := mysql.NewDatabase(database.Config{
		Host:     "127.0.0.1",
		Port:     port,
		User:     username,
		Password: password,
		DbName:   dbName,
	})
	if err != nil {
		log.Fatal(err)
	}
	sqldef.Run(schema.GeneratorModeMysql, database, sp, options)
}

func TestMain(m *testing.M) {
	// DBの立ち上げ
	resource, pool := createContainer()
	// defer closeContainer(resource, pool)

	// DBへ接続する
	db := connectDB(resource, pool)
	defer db.Close()

	// テスト用DBをセットアップ
	setupTestDB()

	// 接続情報を渡すとQueryインスタンスを生成
	query = dbgen.New(db)

	// テスト実行
	m.Run()
}
