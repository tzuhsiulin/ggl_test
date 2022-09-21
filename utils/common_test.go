package utils

import (
	"os"
	"reflect"
	"testing"

	mock_validator "ggl_test/mocks/validator"
	"github.com/go-playground/validator/v10"
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

func TestConvertBindingErr(t *testing.T) {
	var tests = []struct {
		input       func() error
		expectedVal string
	}{
		{
			func() error {
				err := mock_validator.MockFieldError{}
				err.SetTag("required")
				err.SetField("test")
				err.SetKind(reflect.String)
				return validator.ValidationErrors{&err}
			},
			"`test` cannot be empty",
		},
		{
			func() error {
				err := mock_validator.MockFieldError{}
				err.SetTag("max")
				err.SetField("test")
				err.SetKind(reflect.Int)
				err.SetParam("10")
				return validator.ValidationErrors{&err}
			},
			"`test` cannot be greater than 10",
		},
		{
			func() error {
				err := mock_validator.MockFieldError{}
				err.SetTag("max")
				err.SetField("test")
				err.SetKind(reflect.String)
				err.SetParam("10")
				return validator.ValidationErrors{&err}
			},
			"`test` length cannot be greater than 10",
		},
		{
			func() error {
				err := mock_validator.MockFieldError{}
				err.SetTag("min")
				err.SetField("test")
				err.SetKind(reflect.Int)
				err.SetParam("10")
				return validator.ValidationErrors{&err}
			},
			"`test` cannot be less than 10",
		},
		{
			func() error {
				err := mock_validator.MockFieldError{}
				err.SetTag("min")
				err.SetField("test")
				err.SetKind(reflect.String)
				err.SetParam("10")
				return validator.ValidationErrors{&err}
			},
			"`test` length cannot be less than 10",
		},
		{
			func() error {
				err := mock_validator.MockFieldError{}
				err.SetTag("test")
				err.SetField("test")
				err.SetKind(reflect.String)
				err.SetParam("10")
				return validator.ValidationErrors{&err}
			},
			"",
		},
	}
	for _, tt := range tests {
		assert.Equal(t, ConvertBindingErr(tt.input()), tt.expectedVal)
	}
}
