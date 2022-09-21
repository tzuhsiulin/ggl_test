package controllers

import (
	"net/http"

	"ggl_test/config"
	"ggl_test/models/dto"
	"ggl_test/models/entity"
	"ggl_test/service"
	"ggl_test/utils"
	"ggl_test/utils/customerror"
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
		log.GetLoggerWithCtx(appCtx).Info("failed to get task list")
		utils.ResponseError(appCtx, err)
		return
	}

	log.GetLoggerWithCtx(appCtx).Info("got task list")
	utils.Response(appCtx, &dto.GetTaskListResp{Result: *taskList})
}

func (s *TaskController) Add(c *gin.Context) {
	appCtx := &dto.AppContext{
		GinContext: c,
	}

	var req dto.CreateTaskReq
	if err := c.ShouldBind(&req); err != nil {
		log.GetLoggerWithCtx(appCtx).Info("invalid params: %v", req)
		utils.ResponseError(appCtx, customerror.NewErr(customerror.ErrorCodeInvalidParam, utils.ConvertBindingErr(err)))
		return
	}

	taskInfo, err := s.taskService.Add(appCtx, &entity.Task{Name: req.Name})
	if err != nil {
		log.GetLoggerWithCtx(appCtx).Info("failed to create task")
		utils.ResponseError(appCtx, err)
		return
	}

	log.GetLoggerWithCtx(appCtx).Info("created task, and return the task info")
	utils.Response(appCtx, &dto.CreateTaskResp{
		Result: *taskInfo,
	}, http.StatusCreated)
}

func (s *TaskController) Update(c *gin.Context) {
	appCtx := &dto.AppContext{
		GinContext: c,
	}

	var req dto.UpdateTaskReq
	if err := c.ShouldBindUri(&req.Path); err != nil {
		log.GetLoggerWithCtx(appCtx).Info(err)
		utils.ResponseError(appCtx, customerror.NewErr(customerror.ErrorCodeInvalidParam))
		return
	}

	if err := c.ShouldBind(&req.Data); err != nil {
		log.GetLoggerWithCtx(appCtx).Info(err)
		utils.ResponseError(appCtx, customerror.NewErr(customerror.ErrorCodeInvalidParam, utils.ConvertBindingErr(err)))
		return
	}

	taskInfo, err := s.taskService.Update(appCtx, req.Path.Id, &entity.Task{Name: req.Data.Name, Status: req.Data.Status})
	if err != nil {
		log.GetLoggerWithCtx(appCtx).Info("failed to update task info")
		utils.ResponseError(appCtx, err)
		return
	}

	utils.Response(appCtx, &dto.UpdateTaskResp{Result: *taskInfo})
}

func (s *TaskController) Delete(c *gin.Context) {
	appCtx := &dto.AppContext{
		GinContext: c,
	}

	var req dto.DeleteTaskReq
	if err := c.ShouldBindUri(&req); err != nil {
		log.GetLoggerWithCtx(appCtx).Info(err)
		utils.ResponseError(appCtx, customerror.NewErr(customerror.ErrorCodeInvalidParam))
		return
	}

	err := s.taskService.Delete(appCtx, req.Id)
	if err != nil {
		log.GetLoggerWithCtx(appCtx).Info("failed to delete task info")
		utils.ResponseError(appCtx, err)
		return
	}

	utils.Response(appCtx, nil)
}
