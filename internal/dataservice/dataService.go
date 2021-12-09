package dataservice

import (
	"wsc2017/app/container"
	"wsc2017/config"
	"wsc2017/internal/dataservice/mongodb"
	"wsc2017/internal/dataservice/pgsql"
	"wsc2017/internal/logger"
	"wsc2017/vendor/go.mongodb.org/mongo-driver/mongo"
)

type DataStoreInterface interface{}

var dataserviceMap = map[string]DataServiceFBInterface{
	config.MONGODB: mongodb.MongoService{},
	config.POSTGRE: pgsql.PgService{},
}

// The builder interface for factory method pattern
// Every factory needs to implement Build method
type DataServiceFBInterface interface {
	Build(container.Container, *config.DatastoreConfig) (DataStoreInterface, error)
}

func getDataServiceFb(code string) DataServiceFBInterface {
	return dataserviceMap[code]
}

func BuildConnection(c *container.ServiceContainer, code string) DataStoreInterface {
	ds := getDataServiceFb(code)
	ds.Build(c, c.Config)
	if value, found := c.Get(code); found {
		logger.SugarLog.Infof("Found database connection in container with key: %s")
		return value.(*mongo.Database)
	}

}
