package main

import (
	"sendMsg/logger"
	_ "sendMsg/model"
	"sendMsg/server/manager"
)

func main() {
	logger.Info("服务器启动")
	m := manager.Get()
	err := m.Init()
	m.Test.Test()
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
