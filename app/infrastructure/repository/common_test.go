package repository

import (
	"database/sql"
	"fmt"

	"github/code-kakitai/code-kakitai/config"
	"github/code-kakitai/code-kakitai/infrastructure/db/dbgen"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var (
	connectionString string
	driver           = "mysql"
	Queries          *dbgen.Queries
	db               *sql.DB
)

func TestMain(m *testing.M) {
	setDB()
	beforeEach()
	code := m.Run()
	afterEach()
	os.Exit(code)
}

func setDB() {
	dbConfig := config.GetConfig().DB
	connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	Queries = dbgen.New(db)
}

func beforeEach() {
	fmt.Println("before each...")
	// ここでマイグレーションしたい
	// _, err := db.Exec("DROP DATABASE IF EXISTS mydb")
	// if err != nil {
	// 	panic(err)
	// }
	// _, err = db.Exec("CREATE DATABASE mydb")
	// if err != nil {
	// 	panic(err)
	// }
	// cmd := exec.Command("goose", "-dir", "../../../migrations", driver, connectionString, "up")
	// _, err = cmd.Output()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	panic(err)
	// }
	// fmt.Println("migration successfully")
}

func afterEach() {
	fmt.Println("after each...")
}
