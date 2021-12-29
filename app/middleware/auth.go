package middleware

import (
	"errors"
	"wsc2017/domain/repository"
	ierrors "wsc2017/internal/errors"
	"wsc2017/internal/logger"

	"github.com/gofiber/fiber/v2"
)

func (md *Middleware) AuthMiddleware(c *fiber.Ctx) error {
	token := c.Query("token")

	if token == "" {
		logger.Log.Info("token is empty")
		return ierrors.SendErrorResponse(c, ierrors.NewUnauthorizedError(errors.New("token is required")))
	}

	// Get user repository
	userRepository := repository.BuildRepository(md.container, *md.container.Config.RepositoryConfig.User).(*repository.UserRepository)

	// Check token
	user, err := userRepository.GetUserFromToken(token)

	if err != nil {
		logger.SugarLog.Infof("Get user from token error: %v", err)
		return ierrors.SendErrorResponse(c, ierrors.NewInvalidLoginError(err))
	}

	c.Locals("user", user)

	return c.Next()
}
