package main

import (
	"sendMsg/db"
	"sendMsg/logger"
	"sendMsg/model"
	_ "sendMsg/model"
)

func main() {
	logger.Info("服务器启动")
	err := db.InitDbModel()
	if err != nil {
		return
	}
	model.GetTestModel().Insert()
	logger.Info("服务器启动成功")
}
