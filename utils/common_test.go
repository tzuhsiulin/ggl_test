package utils

import (
	"testing"

	"gotest.tools/assert"
)

func TestIsProd(t *testing.T) {
	var tests = []struct {
		input       string
		expectedVal bool
	}{
		{"prod", true},
		{"dev", false},
	}
	for _, tt := range tests {
		assert.Equal(t, IsProd(tt.input), tt.expectedVal)
	}
}
