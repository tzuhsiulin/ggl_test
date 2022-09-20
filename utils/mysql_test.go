package utils

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDataNotFoundErr(t *testing.T) {
	assert.Equal(t, IsDataNotFoundErr(sql.ErrNoRows), true)
	assert.Equal(t, IsDataNotFoundErr(errors.New("tests")), false)
}
