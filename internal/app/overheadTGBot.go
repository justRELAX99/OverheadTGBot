package app

import (
	"OverheadTGBot/internal/logic"
	"OverheadTGBot/internal/repository"
	"OverheadTGBot/pkg/config"
	"OverheadTGBot/pkg/logger"
	"OverheadTGBot/pkg/sqlite"
	"OverheadTGBot/pkg/telego"
)

func Run(configSettings config.Config) {

	//connections

	logger := logger.NewZapLogger(configSettings.Logger)
	sqliteClient := sqlite.NewClient(configSettings.Sqlite, logger)
	telegoClient := telego.NewTelegoClient(configSettings.TelegramBot)
	logger.Info("Connections successfully loaded")

	//repository
	messageRepository := repository.NewMessageRepository(sqliteClient)
	userRepository := repository.NewUserRepository(sqliteClient)

	//logic
	messageLogic := logic.NewMessageLogic(messageRepository)
	userLogic := logic.NewUserLogic(userRepository)

	senderLogic := logic.NewMessageSender(telegoClient, messageLogic)
	receiverLogic := logic.NewSecretReceiver(telegoClient, messageLogic, userLogic, senderLogic)
	logger.Info("Telegram bor ready for work")

	receiverLogic.ReceiveSecret()

	select {}
}
