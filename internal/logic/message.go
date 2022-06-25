package logic

import "OverheadTGBot/internal/model"

type messageLogic struct {
	messageRepository model.MessageRepository
	bot               model.TelegramBot
}

func NewMessageLogic(messageRepository model.MessageRepository, bot model.TelegramBot) model.MessageLogic {
	return messageLogic{
		messageRepository: messageRepository,
		bot:               bot,
	}
}
