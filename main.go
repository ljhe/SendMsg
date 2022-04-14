package main

import (
	"net/http"
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
	http.HandleFunc("/", sayHello)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Err("服务器启动失败:%v", err)
		return
	}
	logger.Info("服务器启动成功")
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	model.GetTestModel().Query()
}
