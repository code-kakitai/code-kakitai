package main

import (
	"github/code-kakitai/code-kakitai/config"
	"github/code-kakitai/code-kakitai/presentation"
)

const port = ":8080"

// @title ~~~ API
// @version バージョン(1.0)
// @description 説明
// @license.name ライセンス(必須)
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
func main() {
	api := config.NewGinEngine()
	presentation.InitRoute(api)
	api.Run(port)
}
