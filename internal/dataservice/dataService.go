package dataservice

import (
	"wsc2017/app/container"
	"wsc2017/config"
)

type DataStoreInterface interface{}

var dataserviceMap = map[string]DataServiceFBInterface{
	config.MONGODB: &MongoService{},
	config.POSTGRE: &PgService{},
}

// The builder interface for factory method pattern
// Every factory needs to implement Build method
type DataServiceFBInterface interface {
	Build(container.Container, *config.DatastoreConfig) (DataStoreInterface, error)
}

func getDataServiceFb(code string) DataServiceFBInterface {
	return dataserviceMap[code]
}

func BuildConnection(c *container.ServiceContainer, dbConfig config.DatastoreConfig) DataStoreInterface {
	code := dbConfig.Code
	dsBuilder := getDataServiceFb(code)

	dsi, err := dsBuilder.Build(c, &dbConfig)

	if err != nil {
		return nil
	}

	return dsi
}
