package main

import (
	"log"
	"os"

	"github.com/yumekumo/sauna-shop/infrastructure/mysql/db/schema"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:   "migration",
			Usage:  "DBマイグレーション",
			Action: schema.Migrate,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
