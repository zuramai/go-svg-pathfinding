package repository

import (
	"context"
	"wsc2017/domain/model"
	"wsc2017/internal/logger"

	"github.com/jackc/pgx/v4"
)

type PlaceRepository struct {
	DB *pgx.Conn
}

func (pr *PlaceRepository) GetAllPlace() ([]model.Place, error) {
	db := pr.DB

	var places []model.Place
	var place model.Place

	rows, err := db.Query(context.Background(), "SELECT * FROM places")
	if err != nil {
		logger.SugarLog.Errorf("Get all place failed %v\n", err)
		return nil, err
	}

	for rows.Next() {

		rows.Scan(&place.Id, &place.Code, &place.Name, &place.Latitude, &place.Longitude, &place.X, &place.Y, &place.ImagePath, &place.Description)

		logger.SugarLog.Debug(place.Id)

		places = append(places, place)
	}

	return places, nil
}

// func (pr *PlaceRepository) GetPlaceByID(id int64) (*model.Place, error) {

// }

// func (pr *PlaceRepository) UpdatePlace(id int64, place *model.Place) (*model.Place, error) {

// }

// func (pr *PlaceRepository) InsertPlace() (*model.Place, error) {

// }

// func (pr *PlaceRepository) DeletePlace() (rowsAffected int, err error) {

// }
