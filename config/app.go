package config

import (
	"os"

	"ggl_test/utils"
)

type AppCfg struct {
	IsProd bool
	DbHost string
	DbUser string
	DbPwd  string
}

func GetAppCfg() *AppCfg {
	isProd := utils.IsProdEnv()
	if isProd {
		return getProdCfg()
	}
	return getDevCfg()
}

func getProdCfg() *AppCfg {
	return &AppCfg{
		IsProd: true,
		DbHost: os.Getenv("DB_HOST"),
		DbUser: os.Getenv("DB_USER"),
		DbPwd:  os.Getenv("DB_PWD"),
	}
}

func getDevCfg() *AppCfg {
	return &AppCfg{
		IsProd: false,
		DbHost: "127.0.0.1",
		DbUser: "root",
		DbPwd:  "test",
	}
}
