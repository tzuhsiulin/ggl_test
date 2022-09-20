package dto

import (
	"ggl_test/models/entity"
	"github.com/gin-gonic/gin"
)

type AppContext struct {
	GinContext *gin.Context
}

type GetTaskListResp struct {
	Result []entity.Task `json:"result"`
}
