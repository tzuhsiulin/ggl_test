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
}
