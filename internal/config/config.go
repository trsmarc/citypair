package config

import (
	"citypair/pkg/log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server *Server `yaml:"server,omitempty"`
}

type Server struct {
	Port string `yaml:"port,omitempty"`
}

func Load(file string, logger log.Logger) (*Config, error) {
	logger.Infof("loading config file : %s \n", file)
	absPath, err := filepath.Abs(file)
	if err != nil {
		return nil, err
	}

	c := Config{}
	rawcfg, err := os.ReadFile(absPath)
	if err != nil {
		logger.Errorf("error reading config file : %s \n", file)
		return nil, err
	}

	if err = yaml.Unmarshal(rawcfg, &c); err != nil {
		logger.Errorf("error on json marshalling of config file : %s \n", file)
		return nil, err
	}

	if os.Getenv("APP_PORT") != "" {
		c.Server.Port = os.Getenv("APP_PORT")
	}

	return &c, nil
}
