package main

import (
	"context"

	"github.com/yumekumo/sauna-shop/config"
	"github.com/yumekumo/sauna-shop/infrastructure/mysql/db"
	"github.com/yumekumo/sauna-shop/infrastructure/redis"
	"github.com/yumekumo/sauna-shop/server"
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
