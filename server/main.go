package main

import (
	"sendMsg/libs/logger"
	_ "sendMsg/server/handler"
	"sendMsg/server/manager"
	_ "sendMsg/server/model"
)

func main() {
	logger.Info("服务器启动")
	m := manager.Get()
	err := m.Init()
	if err != nil {
		logger.Err("服务器初始化失败:err:%v", err)
		return
	}
	err = m.Start()
	if err != nil {
		logger.Err("服务器启动失败:err:%v", err)
		return
	}
	m.Run()
	logger.Info("服务器启动成功")
	manager.WaitForTerminate()
	m.Stop()
	logger.Info("服务器关闭")
}
