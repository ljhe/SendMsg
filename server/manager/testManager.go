package manager

import (
	"sendMsg/libs/logger"
	"sendMsg/libs/manager"
)

type TestManager struct {
	manager.DefaultModule
}

func NewTestManager() *TestManager {
	return &TestManager{}
}

func (t *TestManager) Init() error {
	return nil
}

func (t *TestManager) Test() {
	logger.Info("这里是测试方法")
	m.TestB.Test()
}
