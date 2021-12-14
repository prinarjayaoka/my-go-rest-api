package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/prinarjayaoka/my-go-rest-api/utils"
)

func main() {
	migrationsPath := fmt.Sprintf("file:%s/my-go-rest-api/migrations", utils.Root)
	dbHost := os.Getenv("MY_GO_DB_HOST")
	dbUsername := os.Getenv("MY_GO_DB_USERNAME")
	dbPasswd := os.Getenv("MY_GO_DB_PASSWD")
	dbUse := "redditgo"

	dsnString := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?multiStatements=true",
		dbUsername,
		dbPasswd,
		dbHost,
		dbUse,
	)

	db, _ := sql.Open("mysql", dsnString)

	driver, _ := mysql.WithInstance(db, &mysql.Config{})

	m, _ := migrate.NewWithDatabaseInstance(migrationsPath, "mysql", driver)

	m.Steps(2)
}
