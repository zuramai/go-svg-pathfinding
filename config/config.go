package config

import (
	"io/ioutil"
	"path/filepath"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

type Config struct {
	AppConfig      ApplicationConfig `yaml:"appConfig"`
	DatabaseConfig DatastoreConfig   `yaml:"databaseConfig"`
	WebConfig      WebConfig         `yaml:"webConfig"`
	ZapConfig      zap.Config        `yaml:"zapConfig"`
}

type ApplicationConfig struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}

type DatastoreConfig struct {
	Connection string `yaml:"connection"`
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
}

type WebConfig struct {
	Port string `yaml:"port"`
}

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
