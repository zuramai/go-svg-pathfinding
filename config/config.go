package config

import (
	"io/ioutil"
	"path/filepath"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

type Config struct {
	AppConfig        ApplicationConfig `yaml:"appConfig"`
	RepositoryConfig RepositoryConfig  `yaml:"repositoryConfig"`
	MongoConfig      DatastoreConfig   `yaml:"mongoConfig"`
	PgConfig         DatastoreConfig   `yaml:"pgConfig"`
	WebConfig        WebConfig         `yaml:"webConfig"`
	ZapConfig        zap.Config        `yaml:"zapConfig"`
}

type ApplicationConfig struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}

type DatastoreConfig struct {
	Code     string `yaml:"code"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"dbName"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type WebConfig struct {
	Port string `yaml:"port"`
}

type RepositoryConfig struct {
	Place    *Repository `yaml:"place"`
	Route    *Repository `yaml:"route"`
	User     *Repository `yaml:"user"`
	Schedule *Repository `yaml:"schedule"`
}

type Repository struct {
	Code           string          `yaml:"code"`
	DatabaseConfig DatastoreConfig `yaml:"databaseConfig"`
}

const (
	MONGODB = "mongodb"
	POSTGRE = "pgsql"
)

func BuildConfig(filename string) (*Config, error) {
	var config Config

	// Check for absolute config filenamee
	abs, err := filepath.Abs(filename)
	if err != nil {
		return nil, errors.Wrap(err, "Error read config file")
	}

	file, err := ioutil.ReadFile(abs)
	if err != nil {
		return nil, errors.Wrap(err, "Error read config file")
	}

	// Unmarshal yaml to struct
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, errors.Wrap(err, "Unmarshal error")
	}

	return &config, nil
}
