package dbTest

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
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
	dbName   = "code_kakitai_test"
	port     int // コンテナ起動時に決定する
)

func CreateContainer() (*dockertest.Resource, *dockertest.Pool) {
	pool, err := dockertest.NewPool("")
	pool.MaxWait = time.Minute * 2
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// Dockerコンテナ起動時の細かいオプションを指定する
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

func CloseContainer(resource *dockertest.Resource, pool *dockertest.Pool) {
	// コンテナの終了
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func ConnectDB(resource *dockertest.Resource, pool *dockertest.Pool) *sql.DB {
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

func SetupTestDB(schemaFilePath string) {
	// マイグレーション
	desiredDDLs, err := sqldef.ReadFile(schemaFilePath)
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
