package mongodb

import (
	"context"
	"fmt"
	"time"
	"wsc2017/app/container"
	"wsc2017/config"
	"wsc2017/internal/dataservice"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoService struct{}

func (m *MongoService) Build(c container.Container, dsc *config.DatastoreConfig) (dataservice.DataStoreInterface, error) {
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

	return client.Database(dsc.DBName), nil
}
