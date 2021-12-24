package controller

import (
	"wsc2017/domain/model"
	"wsc2017/domain/usecase"
	"wsc2017/internal/errors"
	"wsc2017/internal/logger"

	"github.com/gofiber/fiber/v2"
)

func (controller *Controller) GetPlace(c *fiber.Ctx) error {
	logger.Log.Info("GetPlace called")

	puc := usecase.GetPlaceUseCase(controller.Container)
	places, err := puc.GetAllPlace()

	if err != nil {
		logger.SugarLog.Infof("Error getplace: %v", err)
		return errors.SendErrorResponse(c, err)
	}

	return places.SendResponse(c)
}

func (controller *Controller) ShowPlace(c *fiber.Ctx) error {
	logger.Log.Info("ShowPlace called")

	puc := usecase.GetPlaceUseCase(controller.Container)

	// parse body
	placeId, err := c.ParamsInt("id")
	if err != nil {
		logger.SugarLog.Errorf("Error parsing params: %v", err)
		return errors.SendErrorResponse(c, err)
	}

	res, err := puc.ShowPlace(int64(placeId))

	if err != nil {
		logger.SugarLog.Infof("Error getplace: %v", err)
		return errors.SendErrorResponse(c, err)
	}

	return res.SendResponse(c)
}
func (controller *Controller) StorePlace(c *fiber.Ctx) error {
	logger.Log.Info("ShowPlace called")

	puc := usecase.GetPlaceUseCase(controller.Container)

	// Parse Request Body
	place := model.Place{}
	err := c.BodyParser(&place)
	if err != nil {
		logger.SugarLog.Errorf("Error parsing params: %v", err)
		return errors.SendErrorResponse(c, err)
	}

	// Parse Image
	image, err := c.FormFile("image")
	if err != nil {
		logger.SugarLog.Errorf("Error parsing image: %v", err)
		return errors.SendErrorResponse(c, err)
	}

	res, err := puc.StorePlace(c, &place, image)

	if err != nil {
		logger.SugarLog.Infof("Error storeplace: %v", err)
		return errors.SendErrorResponse(c, err)
	}

	return res.SendResponse(c)
}
func (controller *Controller) UpdatePlace(c *fiber.Ctx) error {
	return nil
}
func (controller *Controller) DeletePlace(c *fiber.Ctx) error {
	return nil
}
