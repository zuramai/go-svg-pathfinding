package container

import "wsc2017/config"

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
	return value, found
}

func (sc *ServiceContainer) Put(code string, value interface{}) {
	sc.FactoryMap[code] = value
}
