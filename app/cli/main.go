package main

import (
	"fmt"
	"github/code-kakitai/code-kakitai/infrastructure/db/schema"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:  "greet",
			Usage: "fight the loneliness!",
			Action: func(*cli.Context) error {
				fmt.Println("Hello friend!")
				return nil
			},
		},
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
