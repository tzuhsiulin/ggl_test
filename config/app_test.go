package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAppCfg(t *testing.T) {
	os.Setenv("ENV", "dev")
	appCfg := GetAppCfg()
	assert.Equal(t, appCfg.IsProd, false)
	assert.Equal(t, appCfg.DbHost, "127.0.0.1")
	assert.Equal(t, appCfg.DbUser, "root")
	assert.Equal(t, appCfg.DbPwd, "test")

	os.Setenv("ENV", "prod")
	os.Setenv("DB_HOST", "test")
	os.Setenv("DB_USER", "test")
	os.Setenv("DB_PWD", "test")
	appCfg = GetAppCfg()
	assert.Equal(t, appCfg.IsProd, true)
	assert.Equal(t, appCfg.DbHost, "test")
	assert.Equal(t, appCfg.DbUser, "test")
	assert.Equal(t, appCfg.DbPwd, "test")
}
