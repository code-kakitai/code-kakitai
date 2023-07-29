package db

import (
	"database/sql"
	"fmt"
	"github/code-kakitai/code-kakitai/config"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db/dbgen"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const maxRetries = 5
const delay = 5 * time.Second

var query *dbgen.Queries

func NewMainDB() *dbgen.Queries {
	if query != nil {
		return query
	}
	db, err := connect()
	if err != nil {
		panic(err)
	}
	query = dbgen.New(db)
	return query
}

// dbに接続する：最大5回リトライする
func connect() (*sql.DB, error) {
	cfg := config.GetConfig().DB
	for i := 0; i < maxRetries; i++ {
		connect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
		fmt.Println(connect)
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
