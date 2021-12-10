package container

import (
	"wsc2017/config"
	"wsc2017/internal/logger"
)

type Container interface {
	Get(code string) (interface{}, bool)
	Put(code string, value interface{})
}

type ServiceContainer struct {
	FactoryMap map[string]interface{}
	Config     *config.Config
}

func (sc *ServiceContainer) Get(code string) (interface{}, bool) {
	value, found := sc.FactoryMap[code]
	if found {
		logger.SugarLog.Debugf("Finding %s and found", code)
	}
	return value, found
}

func (sc *ServiceContainer) Put(code string, value interface{}) {
	logger.SugarLog.Debugf("Putting %s in container", code)
	sc.FactoryMap[code] = value
}
