package logic

import (
	"OverheadTGBot/internal/model"
	"OverheadTGBot/pkg/errors"
)

type messageLogic struct {
	messageRepository model.MessageRepository
}

func NewMessageLogic(messageRepository model.MessageRepository) model.MessageLogic {
	return messageLogic{
		messageRepository: messageRepository,
	}
}

func (m messageLogic) SaveMessage(messages []model.Message) (err error) {
	err = m.messageRepository.SaveMessage(messages)
	if err != nil {
		return errors.Wrap(err, "Cant save messages")
	}
	return nil
}
