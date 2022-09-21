package server

import (
	"ggl_test/config"
	"ggl_test/controllers"
	"ggl_test/db"
	"ggl_test/models/dto"
	"ggl_test/models/repo"
	"ggl_test/service"
	"ggl_test/utils"
	"ggl_test/utils/customerror"
	"ggl_test/utils/log"
	"github.com/gin-gonic/gin"
)

func AddRoutes(appCfg *config.AppCfg, router *gin.Engine) {
	log.GetLogger().Info("add routes")

	db := db.GetMysqlConn(appCfg)
	taskRepo := repo.NewTaskRepo(db)
	taskSvc := service.NewTaskService(taskRepo)

	router.NoRoute(func(c *gin.Context) {
		utils.ResponseError(&dto.AppContext{GinContext: c}, customerror.NewErr(customerror.ErrorCodeNotFound))
	})

	taskController := controllers.NewTaskController(appCfg, taskSvc)
	router.GET("/tasks", taskController.GetTaskList)
	router.POST("/task", taskController.Add)
	router.PUT("/task/:id", taskController.Update)
	router.DELETE("/task/:id", taskController.Delete)
}
