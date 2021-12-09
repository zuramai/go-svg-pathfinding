package controller

import (
	"wsc2017/domain/usecase"
	"wsc2017/internal/logger"

	"github.com/gofiber/fiber/v2"
)

func (controller *Controller) GetPlace(c *fiber.Ctx) error {
	logger.Log.Info("GetPlace called")

	puc := usecase.GetPlaceUseCase(controller.Container)
	places, err := puc.GetAllPlace()

	if err != nil {
		logger.SugarLog.Infof("Error getplace: %v", err)
	}

	return c.JSON(places)
}
func (controller *Controller) ShowPlace(c *fiber.Ctx) error {
	return nil
}
func (controller *Controller) StorePlace(c *fiber.Ctx) error {
	return nil
}
func (controller *Controller) UpdatePlace(c *fiber.Ctx) error {
	return nil
}
func (controller *Controller) DeletePlace(c *fiber.Ctx) error {
	return nil
}
