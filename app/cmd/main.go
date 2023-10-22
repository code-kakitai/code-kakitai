package main

import (
	"context"

	"github/code-kakitai/code-kakitai/config"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db"
	"github/code-kakitai/code-kakitai/infrastructure/redis"
	"github/code-kakitai/code-kakitai/server"
)

// @title コードカキタイ
// @version バージョン(1.0)
// @description 説明
// @host localhost:8080
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf := config.GetConfig()
	db.NewMainDB(conf.DB)
	db.NewReadDB(conf.ReadDB)

	redisCli := redis.NewClient(conf.Redis)
	defer redisCli.Close()

	server.Run(ctx, conf)
}
