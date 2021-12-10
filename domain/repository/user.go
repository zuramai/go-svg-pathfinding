package repository

import (
	"wsc2017/domain/model"

	"github.com/jackc/pgx/v4"
)

type UserDataInterface interface {
	FindByID(id int64)
	FindByCredentials(username string, password string)
}

type UserRepository struct {
	DB *pgx.Conn
}

func (ur *UserRepository) FindByID(id int64) (*model.User, error) {
	return nil, nil
}

func (ur *UserRepository) FindByCredentials(id int64) (*model.User, error) {
	return nil, nil
}
