package db

import (
	"database/sql"
	"fmt"
	"github/code-kakitai/code-kakitai/config"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db/dbgen"
)

var query *dbgen.Queries

func NewMainDB() *dbgen.Queries {
	if query != nil {
		return query
	}
	cfg := config.GetConfig().DB
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name))
	if err != nil {
		panic(err)
	}
	query = dbgen.New(db)
	return query
}
