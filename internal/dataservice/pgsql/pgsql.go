package pgsql

import (
	"context"
	"fmt"
	"os"
	"wsc2017/app/container"
	"wsc2017/config"
	"wsc2017/internal/dataservice"
	"wsc2017/internal/logger"

	"github.com/jackc/pgx/v4"
)

type PgService struct{}

func (pg *PgService) Build(c container.Container, dsc *config.DatastoreConfig) (dataservice.DataStoreInterface, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", dsc.Username, dsc.Password, dsc.Host, dsc.Port, dsc.DBName)

	conn, err := pgx.Connect(context.Background(), connString)

	logger.SugarLog.Debugf("Connecting to postgresql")

	if err != nil {
		logger.SugarLog.Errorf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	c.Put(dsc.Code, conn)

	defer conn.Close(context.Background())

	return conn, nil
}
