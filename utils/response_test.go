package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"ggl_test/models/dto"
	"ggl_test/utils/customerror"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestResponse(t *testing.T) {
	w := httptest.NewRecorder()
	ginCtx, _ := gin.CreateTestContext(w)
	appCtx := &dto.AppContext{
		GinContext: ginCtx,
	}
	assert.NotPanics(t, func() { Response(appCtx, gin.H{"test": 1}) })
	assert.NotPanics(t, func() { Response(appCtx, gin.H{"test": 1}, http.StatusCreated) })
	assert.NotPanics(t, func() { Response(appCtx, nil) })
}

func TestResponseError(t *testing.T) {
	tests := []struct {
		input *customerror.CustomError
	}{
		{customerror.NewErr(customerror.ErrorCodeInvalidParam)},
		{customerror.NewErr(customerror.ErrorCodeUnknown)},
		{customerror.NewErr(4009)},
		{customerror.NewErr(5009)},
	}
	for _, tt := range tests {
		w := httptest.NewRecorder()
		ginCtx, _ := gin.CreateTestContext(w)
		appCtx := &dto.AppContext{
			GinContext: ginCtx,
		}
		assert.NotPanics(t, func() { ResponseError(appCtx, tt.input) })
	}
}
