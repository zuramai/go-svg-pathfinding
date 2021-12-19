package controller

import (
	"wsc2017/domain/model"
	"wsc2017/domain/usecase"
	"wsc2017/internal/errors"
	"wsc2017/internal/logger"

	"github.com/gofiber/fiber/v2"
)

func (controller *Controller) Login(c *fiber.Ctx) error {
	logger.Log.Info("Login called")

	auc := usecase.GetAuthUseCase(controller.Container)

	var credentials model.User

	// Parse payload request
	if err := c.BodyParser(&credentials); err != nil {
		logger.SugarLog.Errorf("Error parsing request: %v", err)
		return errors.SendErrorResponse(c, err)
	}

	login, err := auc.Login(&credentials)

	if err != nil {
		logger.SugarLog.Errorf("Error login: %v", err)
		return errors.SendErrorResponse(c, err)
	}

	return login.SendResponse(c)
}

func (controller *Controller) Logout(c *fiber.Ctx) error {
	return nil
}
