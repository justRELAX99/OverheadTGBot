package logic

import (
	"OverheadTGBot/internal/entity"
	"OverheadTGBot/pkg/logger"
	"context"
)

const (
	secretCommand = "secret"
)

type secretReceiver struct {
	secretChannel chan entity.Parcel
	messageLogic  entity.MessageLogic
	userLogic     entity.UserLogic
	sender        entity.MessageSender
}

func NewSecretReceiver(
	telegramClient entity.TelegramClient,
	messageLogic entity.MessageLogic,
	userLogic entity.UserLogic,
	sender entity.MessageSender,
) entity.SecretReceiver {
	return secretReceiver{
		secretChannel: telegramClient.HandleCommand(secretCommand),
		messageLogic:  messageLogic,
		userLogic:     userLogic,
		sender:        sender,
	}
}

func (p secretReceiver) ReceiveSecret() {
	go p.receiveSecret()
}

func (p secretReceiver) receiveSecret() {
	ctx := context.Background()
	log := logger.Get()
	var err error
	for secretMessage := range p.secretChannel {
		secretMessage.Message.Status = entity.StatusModerated

		secretMessage.Message.Id, err = p.messageLogic.SaveMessage(ctx, secretMessage.Message)
		if err != nil {
			log.Errorf("Cant save messages,because %v", err.Error())
		}

		secretMessage.Sender.Role = entity.AuthorRole
		err = p.userLogic.SaveUser(ctx, secretMessage.Sender)
		if err != nil {
			log.Errorf("Cant save users,because %v", err.Error())
		}

		p.sender.SendMessage(secretMessage.Message)
	}
}
