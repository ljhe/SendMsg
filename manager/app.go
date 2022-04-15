package manager

import (
	"fmt"
	"runtime/debug"
	"sendMsg/logger"
	"strings"
	"sync"
	"time"
)

type Module interface {
	Init() error
	Start() error
	Run()
	Stop()
	MuCheck() error
}

type DefaultModule struct {
}

func (d DefaultModule) Init() error {
	return nil
}

func (d DefaultModule) Start() error {
	return nil
}

func (d DefaultModule) Run() {
}

func (d DefaultModule) Stop() {
}

func (d DefaultModule) MuCheck() error {
	return nil
}

type DefaultModuleManager struct {
	Module
	Modules []Module
}

func NewDefaultModelManager() *DefaultModuleManager {
	return &DefaultModuleManager{
		Modules: make([]Module, 0),
	}
}

func (d *DefaultModuleManager) Init() error {
	for i := 0; i < len(d.Modules); i++ {
		clsName := fmt.Sprintf("%T", d.Modules[i])
		dotIndex := strings.Index(clsName, ".") + 1
		logger.Info("app|Init :%v", clsName[dotIndex:len(clsName)])
		err := d.Modules[i].Init()
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *DefaultModuleManager) Start() error {
	for i := 0; i < len(d.Modules); i++ {
		clsName := fmt.Sprintf("%T", d.Modules[i])
		dotIndex := strings.Index(clsName, ".") + 1
		logger.Info("app|Start :%v", clsName[dotIndex:len(clsName)])
		err := d.Modules[i].Start()
		if err != nil {
			return err
		}
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Err("app|Start time:%v err:%v panic:%v", time.Now(), err, string(debug.Stack()))
			}
		}()
		ticker := time.NewTicker(5 * time.Minute)
		for {
			select {
			case <-ticker.C:
				logger.Info("muCheck start")
				for i := 0; i < len(d.Modules); i++ {
					clsName := fmt.Sprintf("%T", d.Modules[i])
					dotIndex := strings.Index(clsName, ".") + 1
					logger.Info(clsName[dotIndex:len(clsName)] + " muCheck start")
					err := d.Modules[i].MuCheck()
					if err != nil {
						logger.Err("muCheck err:%v", err)
						return
					}
					logger.Info("%v muCheck stop", clsName[dotIndex:len(clsName)])
				}
				logger.Info("muCheck stop")
			}
		}
	}()

	return nil
}

func (d *DefaultModuleManager) Run() {
	for i := 0; i < len(d.Modules); i++ {
		d.Modules[i].Run()
	}
}

func (d *DefaultModuleManager) Stop() {
	var wg sync.WaitGroup
	for i := 0; i < len(d.Modules); i++ {
		wg.Add(1)
		go func(module Module) {
			module.Stop()
			wg.Done()
		}(d.Modules[i])
	}
	wg.Wait()
}

func (d DefaultModuleManager) MuCheck() error {
	return nil
}

func (d *DefaultModuleManager) AppendModel(module Module) Module {
	d.Modules = append(d.Modules, module)
	return module
}
