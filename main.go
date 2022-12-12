package main

import (
	"ginchat/pkg/global/log"
	"ginchat/router"
	"ginchat/server"
	"ginchat/utils"
	"time"
)

func main() {

	utils.InitConfig()
	log.InitLogger(utils.Conf.Log.Path, utils.Conf.Log.Level)
	utils.InitMysql()
	utils.InitRedis()
	utils.InitSN()
	InitTask()
	// 启动服务
	go server.MyServer.Start()
	c := router.Router()
	err := c.Run(":8081")
	if err != nil {
		return
	}
}

func InitTask() {
	// 定时清理过期链接
	utils.DoTask(time.Duration(utils.Conf.Timeout.DelayHeartbeat)*time.Second,
		time.Duration(utils.Conf.Timeout.HeartbeatHz)*time.Second,
		server.ClearConn, "")
}
