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
	logger.SugarLog.Debugf("Finding %s and found %v", code, value)
	return value, found
}

func (sc *ServiceContainer) Put(code string, value interface{}) {
	logger.SugarLog.Debugf("Putting %s ", code, value)
	sc.FactoryMap[code] = value
}
