package controllers

import (
	"fmt"
	"ginTemp/apps/demo/appctx"
	"ginTemp/ebyte/logger"
	"ginTemp/ebyte/response"
	"github.com/gin-gonic/gin"
)

type Default struct{}

func (con Default) Index(c *gin.Context) {
	fmt.Println(appctx.AppConfig.DEMOAPIServer.Port)
	logger.Info("A message in controller")
	response.Success(c, "success")
}
