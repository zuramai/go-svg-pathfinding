package placerepo

import (
	"wsc2017/domain/model"
	"wsc2017/domain/repository"
)

type PlaceUseCase struct {
	PlaceRepository *repository.PlaceRepository
}

func (puc *PlaceUseCase) GetAllPlace() ([]model.Place, error) {
	places, err := puc.PlaceRepository.GetAllPlace()

	if err != nil {
		return nil, err
	}

	return places, nil
}

func (puc *PlaceUseCase) ShowPlace(id int64) (*model.Place, error) {
	return nil, nil
}
func (puc *PlaceUseCase) StorePlace(place *model.Place) (*model.Place, error) {
	return nil, nil
}
func (puc *PlaceUseCase) UpdatePlace(id int64, place *model.Place) (*model.Place, error) {
	return nil, nil
}
func (puc *PlaceUseCase) DeletePlace(id int64) (rowsAffected int, err error) {
	return 0, nil
}
