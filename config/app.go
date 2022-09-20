package config

import "ggl_test/utils"

type AppCfg struct {
	IsProd bool
}

func GetAppCfg() *AppCfg {
	return &AppCfg{
		IsProd: utils.IsProdEnv(),
	}
}
