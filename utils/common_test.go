package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsProdEnv(t *testing.T) {
	var tests = []struct {
		input       string
		expectedVal bool
	}{
		{"prod", true},
		{"dev", false},
	}
	for _, tt := range tests {
		os.Setenv("ENV", tt.input)
		assert.Equal(t, IsProdEnv(), tt.expectedVal)
	}
}
