package placerepo

import (
	"fmt"
	"mime/multipart"
	"strings"
	"wsc2017/domain/model"
	"wsc2017/domain/repository"
	"wsc2017/internal/common"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type PlaceUseCase struct {
	PlaceRepository *repository.PlaceRepository
}

func (puc *PlaceUseCase) GetAllPlace() (*common.Response, error) {
	places, err := puc.PlaceRepository.GetAllPlace()

	if err != nil {
		return nil, err
	}

	resp := common.NewResponse(200, fiber.Map{"data": places})

	return &resp, nil
}

func (puc *PlaceUseCase) ShowPlace(id int64) (*common.Response, error) {
	places, err := puc.PlaceRepository.GetPlaceById(id)

	if err != nil {
		return nil, err
	}

	resp := common.NewResponse(200, fiber.Map{"data": places})

	return &resp, nil
}
func (puc *PlaceUseCase) StorePlace(c *fiber.Ctx, place *model.Place, image *multipart.FileHeader) (*common.Response, error) {
	imagename := image.Filename

	err := c.SaveFile(image, fmt.Sprintf("./storage/images/places/%s", imagename))
	if err != nil {
		return nil, errors.Wrap(err, "failed to save image")
	}

	place.ImagePath = fmt.Sprintf("%s/storage/images/places/%s", c.BaseURL(), imagename)

	place.Code = strings.ToLower(place.Name)[0:3]

	_, err = puc.PlaceRepository.InsertPlace(place)

	if err != nil {
		return nil, errors.Wrap(err, "failed to insert place")
	}

	resp := common.NewResponse(200, fiber.Map{"message": "create success"})

	return &resp, nil
}
func (puc *PlaceUseCase) UpdatePlace(id int64, place *model.Place) (*common.Response, error) {
	return nil, nil
}
func (puc *PlaceUseCase) DeletePlace(id int64) (rowsAffected int, err error) {
	return 0, nil
}
