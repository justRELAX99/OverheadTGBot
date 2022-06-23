package app

import (
	"OverheadTGBot/pkg/config"
	"OverheadTGBot/pkg/logger"
	"OverheadTGBot/pkg/telego"
)

func Run(configsDirectory string) {
	config := config.LoadConfigSettings(configsDirectory)
	logger := logger.NewZapLogger(config.Logger)
	logger.Info("Configuration and logger successfully loaded")
	botClient := telego.NewClient(config.TelegramBot)
	botClient.RegisterMessageHandler()
	logger.Info("Telegram bor ready for work")
}
