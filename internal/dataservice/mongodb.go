package dataservice

import (
	"context"
	"fmt"
	"time"
	"wsc2017/app/container"
	"wsc2017/config"
	"wsc2017/internal/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoService struct{}

func (m *MongoService) Build(c container.Container, dsc *config.DatastoreConfig) (DataStoreInterface, error) {

	// Check if connection is cached in container
	if value, found := c.Get(dsc.Code); found {
		logger.SugarLog.Debug("Found database connection in container key: %s", dsc.Code)
		return value, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	connString := fmt.Sprintf("mongodb://%s:%d", dsc.Host, dsc.Port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connString))

	// Disconnect the connection
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	db := client.Database(dsc.DBName)

	c.Put(dsc.Code, db)
	c.Put("mongoCtx", ctx)

	return db, nil
}
