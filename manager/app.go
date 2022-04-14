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
	logger.Info("这里是测试moduleLen:%d", len(d.Modules))
	return nil
}

func (d *DefaultModuleManager) Start() error {
	for i := 0; i < len(d.Modules); i++ {
		err := d.Modules[i].Start()
		if err != nil {
			return err
		}
	}

	logger.Info("这里是调用了Start")
	go func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Err("app|Start time:%v err:%v panic:%v", time.Now(), err, string(debug.Stack()))
			}
		}()
		ticker := time.NewTicker(5 * time.Second)
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
	logger.Debug("这里是测试")
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

func (d *DefaultModuleManager) AppendModel(module Module) {
	d.Modules = append(d.Modules, module)
}
