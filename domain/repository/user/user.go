package user

import "wsc2017/domain/model"

type UserDataInterface interface {
	FindByID(id int64)
	FindByCredentials(username string, password string)
}

type UserRepository struct{}

func (ur *UserRepository) FindByID(id int64) (*model.User, error) {
	return nil, nil
}

func (ur *UserRepository) FindByCredentials(id int64) (*model.User, error) {
	return nil, nil
}
