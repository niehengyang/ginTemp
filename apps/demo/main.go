package main

import (
	"ginTemp/apps/demo/appctx"
	"ginTemp/apps/demo/route"
	"ginTemp/ebyte"
	"ginTemp/ebyte/apiserver"
	"ginTemp/ebyte/configure"
	"ginTemp/ebyte/logger"
	"ginTemp/etc"
)

// 语音识别系统
func main() {

	var config etc.Config
	err := configure.Load("../../etc/bussiness.yml", &config)
	if err != nil {
		panic(any("Load system config file error: " + err.Error()))
	}

	config.EByteConfig.Logger.AccessFile = "../../logs/api.access.log"
	config.EByteConfig.Logger.ErrorFile = "../../logs/api.access.log"

	err = ebyte.New(config.EByteConfig)
	if err != nil {
		panic(any("Init EByte framework error:" + err.Error()))
	}

	appctx.AppConfig = config

	asrApiServer := apiserver.NewApiServer(config.DEMOAPIServer)
	asrApiServer.BindRouter(route.Routes)
	err = asrApiServer.Start()
	if err != nil {
		logger.Error("apiServer Start error :" + err.Error())
	}

}
