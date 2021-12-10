package dataservice

import (
	"context"
	"fmt"
	"wsc2017/app/container"
	"wsc2017/config"
	"wsc2017/internal/logger"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

type PgService struct{}

func (pg *PgService) Build(c container.Container, dsc *config.DatastoreConfig) (DataStoreInterface, error) {

	if value, found := c.Get(dsc.Code); found {
		logger.SugarLog.Infof("Found database connection in container with key: %s", dsc.Code)
		return value.(*pgx.Conn), nil
	}

	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", dsc.Username, dsc.Password, dsc.Host, dsc.Port, dsc.DBName)

	conn, err := pgx.Connect(context.Background(), connString)

	logger.SugarLog.Debugf("Connecting to postgresql")

	if err != nil {
		logger.SugarLog.Errorf("Unable to connect to database: %v\n", err)
		return nil, errors.New("Error connecting to postgresql service")
	}

	c.Put(dsc.Code, conn)

	return conn, nil
}
