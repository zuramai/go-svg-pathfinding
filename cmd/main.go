package main

import "wsc2017/app"

const (
	DEV_CONFIG  string = "./config/app.dev.yaml"
	PROD_CONFIG string = "./config/app.prod.yaml"
)

func main() {
	config := DEV_CONFIG

	container, err := app.InitApp(config)

}
