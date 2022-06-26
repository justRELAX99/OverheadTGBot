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
	sqliteClient := sqlite.NewClient(config.Sqlite, logger)
	telegoClient := telego.NewTelegoClient(config.TelegramBot)
	logger.Info("Connections successfully loaded")

	//repository
	messageRepository := repository.NewMessageRepository(sqliteClient)
	userRepository := repository.NewUserRepository(sqliteClient)
	//logic
	messageLogic := logic.NewMessageLogic(messageRepository)
	userLogic := logic.NewUserLogic(userRepository)

	brokerLogic := logic.NewBrokerLogic(telegoClient, telegoClient, messageLogic, userLogic)
	logger.Info("Telegram bor ready for work")

	brokerLogic.RedirectParcels()

}
