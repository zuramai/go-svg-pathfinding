package app

import (
	"wsc2017/app/container"
	"wsc2017/config"
	"wsc2017/internal/logger"

	"github.com/pkg/errors"

	"go.uber.org/zap"
)

func InitApp(filename string) (c container.Container, err error) {
	conf, err := config.BuildConfig(filename)

	if err != nil {
		return nil, errors.Wrap(err, "Error building config")
	}

	err = initLogger(conf.ZapConfig)
	if err != nil {
		return nil, err
	}

	return initContainer(conf)
}

func initContainer(conf *config.Config) (c container.Container, err error) {
	factoryMap := make(map[string]interface{})
	c = &container.ServiceContainer{FactoryMap: factoryMap, Config: conf}
	return c, nil
}

func initLogger(zc zap.Config) error {
	err := logger.SetLogger(zc)
	if err != nil {
		return err
	}

	logger.Log.Debug("Logger construction succeed")
	return nil
}
