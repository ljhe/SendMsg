package manager

import (
	"sendMsg/libs/logger"
	"sendMsg/libs/manager"
)

type TestBManager struct {
	manager.DefaultModule
}

func NewTestBManager() *TestBManager {
	return &TestBManager{}
}

func (t *TestBManager) Init() error {
	return nil
}

func (t *TestBManager) Test() {
	logger.Info("这里是测试方法")
}
