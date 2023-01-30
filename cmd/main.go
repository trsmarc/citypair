package main

import (
	"context"
	"flag"

	logger "log"
	"os"

	"citypair/cmd/server"
	"citypair/internal/config"
	"citypair/pkg/log"
)

var Version = "1.0.0"
var configPath = flag.String("config", "config/local.yaml", "path to the config file")

func main() {
	if err := run(); err != nil {
		logger.Println("error :", err)
		os.Exit(1)
	}
}

func initialize() (*config.Config, log.Logger, error) {
	logger := log.New().With(context.Background(), "version", Version)
	flag.Parse()

	cfg, err := config.Load(*configPath, logger)
	if err != nil {
		logger.Errorf("failed to load application configuration: %s", err)
		return nil, nil, err
	}

	return cfg, logger, nil
}

func run() error {
	cfg, logger, err := initialize()
	if err != nil {
		logger.Errorf("failed to initialize: %s", err)
		return err
	}

	r := server.Routing(logger)
	return server.Start(cfg, r, logger)
}
