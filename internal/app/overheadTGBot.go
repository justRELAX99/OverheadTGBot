package app

import (
	"OverheadTGBot/internal/logic"
	"OverheadTGBot/internal/repository"
	"OverheadTGBot/pkg/config"
	"OverheadTGBot/pkg/logger"
	"OverheadTGBot/pkg/sqlite"
	"OverheadTGBot/pkg/telego"
)

func Run(configsDirectory string) {

	//connections
	config := config.LoadConfigSettings(configsDirectory)
	logger := logger.NewZapLogger(config.Logger)
	botClient := telego.NewClient(config.TelegramBot)
	sqliteClient := sqlite.NewClient(config.Sqlite, logger)

	logger.Info("Connections successfully loaded")

	//repository
	messageRepository := repository.NewMessageRepository(sqliteClient)

	//logic
	messageLogic := logic.NewMessageLogic(messageRepository, botClient)

	botClient.RegisterMessageHandler()
	logger.Info("Telegram bor ready for work")
}
