package manager

import (
	"os"
	"os/signal"
	"sendMsg/db"
	"sendMsg/manager"
	"syscall"
)

var m = &ModuleManager{
	DefaultModuleManager: manager.NewDefaultModelManager(),
}

func Get() *ModuleManager {
	return m
}

func (m *ModuleManager) Init() error {
	err := db.InitDbModel()
	m.init()
	err = m.DefaultModuleManager.Init()
	if err != nil {
		return err
	}
	return err
}

func (m *ModuleManager) Start() error {
	err := m.DefaultModuleManager.Start()
	return err
}

func (m *ModuleManager) Run() {
	m.DefaultModuleManager.Run()
}

func (m *ModuleManager) Stop() {
	m.DefaultModuleManager.Stop()
}

// WaitForTerminate wait signal to end the program
func WaitForTerminate() {
	exitChan := make(chan struct{})
	signalChan := make(chan os.Signal, 1)
	go func() {
		<-signalChan
		close(exitChan)
	}()
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-exitChan
}
