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

func (m messageLogic) SaveMessagesForModerate(ctx context.Context, messages []entity.Message) (err error) {
	if len(messages) == 0 {
		return nil
	}
	for _, message := range messages {
		err = m.messageRepository.SaveMessageForModerate(ctx, message)
		if err != nil {
			return errors.Wrap(err, "cant save messages")
		}
	}
	return nil
}

func (m messageLogic) SaveMessageForModerate(ctx context.Context, message entity.Message) (err error) {
	err = m.messageRepository.SaveMessageForModerate(ctx, message)
	if err != nil {
		return errors.Wrap(err, "cant save message")
	}
	return nil
}
