// Package usecase defines all the interfaces for a Micro-service application.
// It is the entry point for the application's business logic. It is a top level package for a Micro-service application.
// This top level package only defines interface, the concrete implementations are defined in sub-package of it.
// It only depends on model package. No other package should dependent on it except cmd.
package usecase

import (
	"wsc2017/app/container"
	"wsc2017/domain/model"
	"wsc2017/domain/repository"
	authUc "wsc2017/domain/usecase/auth"
	placeUc "wsc2017/domain/usecase/place"
	"wsc2017/internal/common"
)

type AuthUseCaseInterface interface {
	Login(credentials *model.User) (*common.Response, error)
	Logout(token *model.LoginToken) (*model.User, error)
}

type PlaceUseCaseInterface interface {
	GetAllPlace() ([]model.Place, error)
	ShowPlace(id int64) (*model.Place, error)
	StorePlace(place *model.Place) (*model.Place, error)
	UpdatePlace(id int64, place *model.Place) (*model.Place, error)
	DeletePlace(id int64) (rowsAffected int, err error)
}

func GetPlaceUseCase(c *container.ServiceContainer) PlaceUseCaseInterface {
	placeRepo := repository.BuildRepository(c, *c.Config.RepositoryConfig.Place)

	return &placeUc.PlaceUseCase{PlaceRepository: placeRepo.(*repository.PlaceRepository)}
}

func GetAuthUseCase(c *container.ServiceContainer) AuthUseCaseInterface {
	userRepo := repository.BuildRepository(c, *c.Config.RepositoryConfig.User)

	return &authUc.AuthUseCase{UserRepository: userRepo.(*repository.UserRepository)}
}
