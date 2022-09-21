package main

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"time"

	"ggl_test/config"
	"ggl_test/utils/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	appCfg := config.GetAppCfg()
	log.GetLogger().Info("get db migration script dir path")
	migrationScriptDirPath, err := filepath.Abs("data/migration_scripts")
	if err != nil {
		panic(err)
	}
	log.GetLogger().Info(migrationScriptDirPath, appCfg)

	log.GetLogger().Info("get db connection")
	db, _ := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:3306)/ggl_test?multiStatements=true",
			appCfg.DbUser, appCfg.DbPwd, appCfg.DbHost))

	retryTimes := 0
	for {
		err := db.Ping()
		if err != nil {
			retryTimes += 1
			log.GetLogger().Info("retry to connect to mysql")
			time.Sleep(time.Second * 5)
			if retryTimes > 5 {
				panic("failed to connect to mysql")
			}
		} else {
			break
		}
	}

	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file:///%s", migrationScriptDirPath),
		"mysql",
		driver,
	)
	if err != nil {
		panic(err)
	}

	log.GetLogger().Info("do db migration")
	m.Up()
}
