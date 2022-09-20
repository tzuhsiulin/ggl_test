package controllers

import (
	"net/http"

	"ggl_test/config"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	appCfg *config.AppCfg
}

func NewTaskController(appCfg *config.AppCfg) *TaskController {
	return &TaskController{appCfg: appCfg}
}

func (s *TaskController) GetTaskList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"test": 123})
}
