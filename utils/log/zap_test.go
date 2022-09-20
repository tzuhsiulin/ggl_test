package log

import (
	"os"
	"reflect"
	"testing"

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
		logger := GetLogger()
		assert.NotNil(t, logger)
		assert.Equal(t, reflect.TypeOf(logger).String(), tt.expectedVal)
	}
}
