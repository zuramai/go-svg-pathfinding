package repository

import (
	"wsc2017/app/container"
	"wsc2017/config"
	"wsc2017/internal/dataservice"
	"wsc2017/internal/logger"

	"github.com/jackc/pgx/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type repositoryFbInterface interface{}

func BuildRepository(c *container.ServiceContainer, repoconf config.Repository) repositoryFbInterface {
	db := dataservice.BuildConnection(c, repoconf.DatabaseConfig)

	switch repoconf.Code {
	case config.PLACE:
		return &PlaceRepository{DB: db.(*pgx.Conn)}
	case config.SCHEDULE:
		return &ScheduleRepository{DB: db.(*pgx.Conn)}
	case config.ROUTE:
		return &RouteRepository{DB: db.(*mongo.Database)}
	case config.USER:
		return &UserRepository{DB: db.(*pgx.Conn)}
	}

	logger.SugarLog.Fatal("Cannot get repository")

	return nil
}
