package main

import (
	"context"
	"github/code-kakitai/code-kakitai/config"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db"
	"github/code-kakitai/code-kakitai/presentation"
	"github/code-kakitai/code-kakitai/presentation/settings"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	api := settings.NewGinEngine()
	query := db.NewMainDB()
	presentation.InitRoute(api,query)

	address := ":" + conf.Server.Port
	srv := &http.Server{
		Addr:              address,
		Handler:           api,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Minute,
		WriteTimeout:      10 * time.Minute,
	}
	go func() {
		// srv.Shutdownが呼ばれるとhttp.ErrServerClosedを返すのでこれは無視する
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		os.Exit(1)
	}
}
