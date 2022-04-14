package manager

import "sendMsg/manager"

type ModuleManager struct {
	*manager.DefaultModuleManager
	Test         *TestManager
	TestBManager *TestBManager
}

func (m *ModuleManager) init() {
	m.AppendModel(NewTestManager())
	m.AppendModel(NewTestBManager())
}
