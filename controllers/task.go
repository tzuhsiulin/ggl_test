package controllers

import (
	"ggl_test/config"
	"ggl_test/models/dto"
	"ggl_test/service"
	"ggl_test/utils"
	"ggl_test/utils/log"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	appCfg      *config.AppCfg
	taskService service.ITaskService
}

func NewTaskController(appCfg *config.AppCfg, taskService service.ITaskService) *TaskController {
	return &TaskController{appCfg: appCfg, taskService: taskService}
}

func (s *TaskController) GetTaskList(c *gin.Context) {
	appCtx := &dto.AppContext{
		GinContext: c,
	}

	log.GetLoggerWithCtx(appCtx).Info("getting task list")
	taskList, err := s.taskService.GetAll(appCtx)
	if err != nil {
		log.GetLoggerWithCtx(appCtx).Error("failed to get task list")
		utils.ResponseError(appCtx, err)
		return
	}

	log.GetLoggerWithCtx(appCtx).Info("got task list")
	utils.Response(appCtx, &dto.GetTaskListResp{Result: *taskList})
}
