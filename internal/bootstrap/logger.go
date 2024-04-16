package bootstrap

import (
	"go_plate/pkg/config"

	"go.uber.org/zap"
)

func NewLogger(cfg *config.Config) (*zap.Logger, error) {
	if cfg.App.Production {
		logger, err := zap.NewProduction()
		return logger, err
	}

	logger, err := zap.NewDevelopment()
	return logger, err
}
