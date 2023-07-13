package schema

import (
	"fmt"
	"github/code-kakitai/code-kakitai/config"
	"log"
	"strconv"

	"github.com/k0kubun/sqldef"
	"github.com/k0kubun/sqldef/database"
	"github.com/k0kubun/sqldef/database/mysql"
	"github.com/k0kubun/sqldef/schema"
	"github.com/urfave/cli"

	"github.com/k0kubun/sqldef/parser"
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
	sqlParser := database.NewParser(parser.ParserModeMysql)
	desiredDDLs, err := sqldef.ReadFile(schemaFile)
	if err != nil {
		return fmt.Errorf("failed to read schema file: %s", err)
	}

	var dryRun bool
	if cCtx.Args().Get(1) == "apply" {
		dryRun = false
	} else {
		dryRun = true
	}
	options := &sqldef.Options{
		DesiredDDLs: desiredDDLs,
		DryRun:      dryRun,
	}

	sqldef.Run(schema.GeneratorModeMysql, db, sqlParser, options)
	return nil
}
