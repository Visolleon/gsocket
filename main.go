package main

import (
	"config"
	"hub"
	"log"
	"utils"
	"web"

	"git.happyseven.cn/golibs/logger"

	"runtime"
)

func init() {
	log.Printf("Server Version: %s;  DEBUG: %s\n", utils.Version, utils.DebugInfo)

	config.Init()
	logger.Init(config.Cfg)
	hub.Init()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// Socket服务启动
	web.Init()
}
