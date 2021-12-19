package auth

import (
	"wsc2017/domain/model"
	"wsc2017/domain/repository"
	"wsc2017/internal/common"
	"wsc2017/internal/errors"

	"github.com/gofiber/fiber/v2"
)

type AuthUseCase struct {
	UserRepository *repository.UserRepository
}

func (auc *AuthUseCase) Login(credentials *model.User) (*common.Response, error) {

	user, err := auc.UserRepository.FindByCredentials(credentials.Username, credentials.Password)

	if err != nil {
		return nil, errors.NewInvalidLoginError(err)
	}

	token, err := auc.UserRepository.CreateToken(user)

	if err != nil {
		return nil, errors.NewInvalidLoginError(err)
	}

	response := common.NewResponse(fiber.Map{
		"token": token,
		"role":  user.Role,
	})

	return &response, nil
}

func (auc *AuthUseCase) Logout(token *model.LoginToken) (*model.User, error) {
	return nil, nil
}
