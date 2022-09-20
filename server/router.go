package server

import (
	"ggl_test/config"
	"ggl_test/controllers"
	"ggl_test/utils/log"
	"github.com/gin-gonic/gin"
)

func AddRoutes(appCfg *config.AppCfg, router *gin.Engine) {
	log.GetLogger().Info("add routes")
	taskController := controllers.NewTaskController(appCfg)
	router.GET("/tasks", taskController.GetTaskList)
}
