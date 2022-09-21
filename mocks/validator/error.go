package mock_validator

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

type MockFieldError struct {
	validator.FieldError
	tag   string
	field string
	kind  reflect.Kind
	param string
}

func (s *MockFieldError) SetTag(tag string) {
	s.tag = tag
}

func (s *MockFieldError) SetField(f string) {
	s.field = f
}

func (s *MockFieldError) SetKind(k reflect.Kind) {
	s.kind = k
}

func (s *MockFieldError) SetParam(p string) {
	s.param = p
}

func (s *MockFieldError) Field() string {
	return s.field
}

func (s *MockFieldError) ActualTag() string {
	return s.tag
}

func (s *MockFieldError) Kind() reflect.Kind {
	return s.kind
}

func (s *MockFieldError) Param() string {
	return s.param
}
