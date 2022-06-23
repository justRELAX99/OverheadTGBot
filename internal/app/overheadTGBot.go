package app

import (
	"OverheadTGBot/pkg/config"
	"OverheadTGBot/pkg/logger"
)

func Run(configsDirectory string) {
	config := config.LoadConfigSettings(configsDirectory)
	logger := logger.NewZapLogger(config.Logger)
	logger.Info("Configuration and logger successfully loaded")
}
