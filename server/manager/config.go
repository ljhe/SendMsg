package manager

import "sendMsg/libs/manager"

type ModuleManager struct {
	*manager.DefaultModuleManager
	Test  *TestManager
	TestB *TestBManager
}

func (m *ModuleManager) init() {
	m.Test = m.AppendModel(NewTestManager()).(*TestManager)
	m.TestB = m.AppendModel(NewTestBManager()).(*TestBManager)
}
