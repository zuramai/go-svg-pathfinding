// Package usecase defines all the interfaces for a Micro-service application.
// It is the entry point for the application's business logic. It is a top level package for a Micro-service application.
// This top level package only defines interface, the concrete implementations are defined in sub-package of it.
// It only depends on model package. No other package should dependent on it except cmd.
package usecase

import (
	"wsc2017/app/container"
	"wsc2017/domain/model"
	"wsc2017/domain/repository"
	placerepo "wsc2017/domain/usecase/place"
	"wsc2017/internal/logger"

	"github.com/jackc/pgx/v4"
)

type AuthUseCaseInterface interface {
	Login()
	Logout()
}

type PlaceUseCaseInterface interface {
	GetAllPlace() ([]*model.Place, error)
	ShowPlace(id int64) (*model.Place, error)
	StorePlace(place *model.Place) (*model.Place, error)
	UpdatePlace(id int64, place *model.Place) (*model.Place, error)
	DeletePlace(id int64) (rowsAffected int, err error)
}

func GetPlaceUseCase(c *container.ServiceContainer) PlaceUseCaseInterface {
	db, found := c.Get(c.Config.RepositoryConfig.Place.DatabaseConfig.Code)
	if found {
		logger.Log.Error("DB NOT FOUND IN CONTAINER")
	}
	logger.SugarLog.Debugf("The content of db %s is: %v", c.Config.PgConfig.Code, db)
	repo := repository.PlaceRepository{DB: db.(*pgx.Conn)}

	return &placerepo.PlaceUseCase{PlaceRepository: &repo}
}
