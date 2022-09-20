package db

import (
	"database/sql"
	"fmt"

	"ggl_test/config"
	_ "github.com/go-sql-driver/mysql"
)

func GetMysqlConn(cfg *config.AppCfg) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/ggl_test?charset=utf8mb4&multiStatements=false&parseTime=True",
		cfg.DbUser, cfg.DbPwd, cfg.DbHost)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return db
}
