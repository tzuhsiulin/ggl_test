package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"ggl_test/models/dto"
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
}
