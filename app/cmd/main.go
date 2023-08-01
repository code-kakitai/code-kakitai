package main

import (
	"context"

	"github/code-kakitai/code-kakitai/config"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db"
	"github/code-kakitai/code-kakitai/server"
)

// @title アプリケーション名
// @version バージョン(1.0)
// @description 説明
// @license.name ライセンス(必須)
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	conf := config.GetConfig()
	db.NewMainDB()
	server.Run(ctx, conf)
}
