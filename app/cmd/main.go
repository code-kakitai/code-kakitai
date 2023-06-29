package main

import (
	"github/code-kakitai/code-kakitai/presentation"
	"github/code-kakitai/code-kakitai/presentation/settings"
)

const port = ":8080"

// @title アプリケーション名
// @version バージョン(1.0)
// @description 説明
// @license.name ライセンス(必須)
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080a
func main() {
	api := settings.NewGinEngine()
	presentation.InitRoute(api)
	api.Run(port)
}
