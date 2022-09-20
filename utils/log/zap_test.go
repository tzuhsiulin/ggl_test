package log

import (
	"os"
	"reflect"
	"testing"

	"ggl_test/models/dto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetLogger(t *testing.T) {
	var tests = []struct {
		input       string
		expectedVal string
	}{
		{"dev", "*zap.SugaredLogger"},
		{"prod", "*zap.SugaredLogger"},
	}
	for _, tt := range tests {
		os.Setenv("ENV", tt.input)
		logger = nil
		logger := GetLogger()
		assert.NotNil(t, logger)
		assert.Equal(t, reflect.TypeOf(logger).String(), tt.expectedVal)
	}
}

func TestGetLoggerWithCtx(t *testing.T) {
	ginCtx := &gin.Context{}
	ctx := &dto.AppContext{GinContext: ginCtx}
	logger := GetLoggerWithCtx(ctx)
	assert.NotNil(t, logger)
	assert.Equal(t, reflect.TypeOf(logger).String(), "*zap.SugaredLogger")
}
