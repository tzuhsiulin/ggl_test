package dto

import (
	"github.com/gin-gonic/gin"
)

type AppContext struct {
	GinContext *gin.Context
}
