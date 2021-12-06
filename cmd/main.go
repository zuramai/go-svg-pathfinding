package main

import (
	"fmt"
	"wsc2017/app"
	"wsc2017/app/container"
	"wsc2017/internal/logger"
	"wsc2017/internal/ports/http"
)

const (
	DEV_CONFIG  string = "./config/app.dev.yaml"
	PROD_CONFIG string = "./config/app.prod.yaml"
)

func main() {
	config := DEV_CONFIG

	c, err := app.InitApp(config)

	if err != nil {
		fmt.Printf("%+v\n", err)
		panic(err)
	}

	sc := c.(*container.ServiceContainer)

	if err := http.RunServer(sc); err != nil {
		logger.SugarLog.Errorf("Failed to run user server: %+v\n", err)
	}
}
