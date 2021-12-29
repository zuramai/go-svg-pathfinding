package repository

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"wsc2017/domain/model"
	"wsc2017/internal/logger"

	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserDataInterface interface {
	FindByID(id int64)
	FindByCredentials(username string, password string)
	CreateToken(username string) string
}

type UserRepository struct {
	DB *pgx.Conn
}

func (ur *UserRepository) FindByID(id int64) (*model.User, error) {
	return nil, nil
}

func (ur *UserRepository) FindByCredentials(username string, password string) (*model.User, error) {
	var user model.User

	err := ur.DB.QueryRow(context.Background(), "SELECT * FROM users WHERE username=$1 LIMIT 1", username).
		Scan(&user.Id, &user.Username, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}
	logger.Log.Sugar().Debugf("Queryrow err %s", user.Password)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) CreateToken(user *model.User) (string, error) {
	logger.Log.Sugar().Infof("Create token called for %s", user.Username)
	// Generate MD5 Hash
	token := md5.Sum([]byte(user.Username))
	tokenString := hex.EncodeToString(token[:])

	// Insert token to database
	_, err := ur.DB.Exec(context.Background(), "INSERT INTO login_tokens(user_id, token) VALUES($1,$2) ON CONFLICT (user_id) DO NOTHING", user.Id, tokenString)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (ur *UserRepository) RevokeToken(token string) error {
	_, err := ur.DB.Exec(context.Background(), "DELETE FROM login_tokens WHERE token=$1", token)

	if err != nil {
		logger.SugarLog.Fatalf("Revoke login token error %v", err)
	}

	return nil
}

func (ur *UserRepository) GetUserFromToken(token string) (*model.User, error) {
	var user model.User

	logger.SugarLog.Info("Searching for token: ", token)

	err := ur.DB.QueryRow(context.Background(), "SELECT users.* FROM users JOIN login_tokens lt ON lt.user_id = users.id WHERE lt.token = $1", token).
		Scan(&user.Id, &user.Username, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
