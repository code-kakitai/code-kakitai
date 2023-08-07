package schema

import (
	"fmt"
	"log"
	"strconv"

	"github.com/k0kubun/sqldef"
	"github.com/k0kubun/sqldef/database"
	"github.com/k0kubun/sqldef/database/mysql"
	"github.com/k0kubun/sqldef/parser"
	"github.com/k0kubun/sqldef/schema"
	"github.com/urfave/cli"

	"github/code-kakitai/code-kakitai/config"
)

// urfave/cli経由で実行する
func Migrate(cCtx *cli.Context) error {
	schemaFile := cCtx.Args().Get(0)
	// データベースへの接続情報を設定
	dbCfg := config.GetConfig().DB

	port, _ := strconv.Atoi(dbCfg.Port)
	db, err := mysql.NewDatabase(database.Config{
		Host:     dbCfg.Host,
		Port:     port,
		User:     dbCfg.User,
		Password: dbCfg.Password,
		DbName:   dbCfg.Name,
	})
	if err != nil {
		log.Fatal(err)
	}
	desiredDDLs, err := sqldef.ReadFile(schemaFile)
	if err != nil {
		return fmt.Errorf("failed to read schema file: %s", err)
	}

	dryRun := true
	if cCtx.Args().Get(1) == "apply" {
		dryRun = false
	}
	options := &sqldef.Options{
		DesiredDDLs:     desiredDDLs,
		DryRun:          dryRun,
		EnableDropTable: true,
	}

	sp := database.NewParser(parser.ParserModeMysql)
	sqldef.Run(schema.GeneratorModeMysql, db, sp, options)
	return nil
}
