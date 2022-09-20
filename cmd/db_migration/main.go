package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"ggl_test/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var dbHost, dbUser, dbPwd string
	if utils.IsProdEnv() {
		dbHost = os.Getenv("DB_HOST")
		dbUser = os.Getenv("DB_USER")
		dbPwd = os.Getenv("DB_PWD")
	} else {
		dbHost = "127.0.0.1"
		dbUser = "root"
		dbPwd = "test"
	}
	migrationScriptDirPath, err := filepath.Abs("data/migration_scripts")
	if err != nil {
		panic(err)
	}

	db, _ := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/ggl_test?multiStatements=true", dbUser, dbPwd, dbHost))
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file:///%s", migrationScriptDirPath),
		"mysql",
		driver,
	)
	if err != nil {
		panic(err)
	}
	m.Up()
}
