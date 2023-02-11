package logic

import (
	"OverheadTGBot/internal/entity"
	"OverheadTGBot/pkg/errors"
	"context"
)

type messageLogic struct {
	messageRepository entity.MessageRepository
}

func NewMessageLogic(messageRepository entity.MessageRepository) entity.MessageLogic {
	return messageLogic{
		messageRepository: messageRepository,
	}
}

func (m messageLogic) SaveMessage(ctx context.Context, message entity.Message) (messageId int, err error) {
	messageId, err = m.messageRepository.SaveMessage(ctx, message)
	if err != nil {
		return messageId, errors.Wrap(err, "cant save message")
	}
	return messageId, nil
}

func (m messageLogic) UpdateStatus(ctx context.Context, messageId int, status string) (err error) {
	if messageId == 0 || status == "" {
		return nil
	}
	err = m.messageRepository.UpdateStatus(ctx, messageId, status)
	if err != nil {
		return errors.Wrapf(err, "cant update status for message with id=%v", messageId)
	}
	return nil
}
