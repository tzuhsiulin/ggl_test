package server

import (
	"net/http"
	"time"

	"ggl_test/config"
	"ggl_test/middlewares"
	"ggl_test/utils/log"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func Start() {
	log.GetLogger().Info("starting server")
	appCfg := config.GetAppCfg()

	if appCfg.IsProd {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()
	router.Use(requestid.New())
	router.Use(gin.Recovery())
	if appCfg.IsProd {
		router.Use(middlewares.JSONLogMiddleware())
	}
	AddRoutes(appCfg, router)
	startListen(appCfg, router)
}

func startListen(appCfg *config.AppCfg, router *gin.Engine) {
	serverConfig := &http.Server{
		Addr:           "127.0.0.1:8080",
		Handler:        router,
		ReadTimeout:    2000 * time.Second,
		WriteTimeout:   2000 * time.Second,
		MaxHeaderBytes: 6000 << 20,
	}
	if appCfg.IsProd {
		serverConfig.Addr = ":8080"
	}
	serverConfig.ListenAndServe()
}
