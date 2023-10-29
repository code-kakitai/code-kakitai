package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github/code-kakitai/code-kakitai/config"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db/dbgen"
)

const maxRetries = 5
const delay = 5 * time.Second

var (
	once      sync.Once
	readOnce  sync.Once
	query     *dbgen.Queries
	readQuery *dbgen.Queries
	dbcon     *sql.DB
	readDBCon *sql.DB
)

// contextからQueriesを取得する。contextにQueriesが存在しない場合は、パッケージ変数からQueriesを取得する
func GetQuery(ctx context.Context) *dbgen.Queries {
	txq := getQueriesWithContext(ctx)
	if txq != nil {
		return txq
	}
	return query
}

func GetReadQuery() *dbgen.Queries {
	return readQuery
}

func SetQuery(q *dbgen.Queries) {
	query = q
}

func SetReadQuery(q *dbgen.Queries) {
	readQuery = q
}

func GetDB() *sql.DB {
	return dbcon
}

func SetDB(d *sql.DB) {
	dbcon = d
}

func SetReadDB(d *sql.DB) {
	readDBCon = d
}

func NewMainDB(cnf config.DBConfig) {
	once.Do(func() {
		dbcon, err := connect(
			cnf.User,
			cnf.Password,
			cnf.Host,
			cnf.Port,
			cnf.Name,
		)
		if err != nil {
			panic(err)
		}
		q := dbgen.New(dbcon)
		SetQuery(q)
		SetDB(dbcon)
	})
}

func NewReadDB(cnf config.ReadDBConfig) {
	readOnce.Do(func() {
		dbcon, err := connect(
			cnf.User,
			cnf.Password,
			cnf.Host,
			cnf.Port,
			cnf.Name,
		)
		if err != nil {
			panic(err)
		}
		q := dbgen.New(dbcon)
		SetReadQuery(q)
		SetReadDB(dbcon)
	})
}

// dbに接続する：最大5回リトライする
func connect(user string, password string, host string, port string, name string) (*sql.DB, error) {
	for i := 0; i < maxRetries; i++ {
		connect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, name)
		db, err := sql.Open("mysql", connect)
		if err != nil {
			return nil, fmt.Errorf("could not open db: %w", err)
		}

		err = db.Ping()
		if err == nil {
			return db, nil
		}

		log.Printf("could not connect to db: %v", err)
		log.Printf("retrying in %v seconds...", delay/time.Second)
		time.Sleep(delay)
	}

	return nil, fmt.Errorf("could not connect to db after %d attempts", maxRetries)
}

type CtxKey string

const (
	QueriesKey CtxKey = "queries"
)

func WithQueries(ctx context.Context, q *dbgen.Queries) context.Context {
	return context.WithValue(ctx, QueriesKey, q)
}

func getQueriesWithContext(ctx context.Context) *dbgen.Queries {
	queries, ok := ctx.Value(QueriesKey).(*dbgen.Queries)
	if !ok {
		return nil
	}
	return queries
}
