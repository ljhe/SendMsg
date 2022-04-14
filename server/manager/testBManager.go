package manager

import (
	"sendMsg/logger"
	"sendMsg/manager"
)

type TestBManager struct {
	manager.DefaultModuleManager
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
