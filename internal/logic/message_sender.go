package logic

import (
	"OverheadTGBot/internal/entity"
	"OverheadTGBot/pkg/logger"
	"context"
)

type messageSender struct {
	messageChannel chan entity.Message
	messageLogic   entity.MessageLogic
	telegramClient entity.TelegramClient
}

func NewMessageSender(
	telegramClient entity.TelegramClient,
	messageLogic entity.MessageLogic,
) entity.MessageSender {
	m := messageSender{
		messageLogic:   messageLogic,
		messageChannel: make(chan entity.Message),
		telegramClient: telegramClient,
	}
	go m.handleSendParcel()
	return m
}

func (m messageSender) SendMessage(message entity.Message) {
	m.messageChannel <- message
}

func (m messageSender) handleSendParcel() {
	ctx := context.Background()
	log := logger.Get()
	for message := range m.messageChannel {
		err := m.telegramClient.SendToAdminChannel(message.Text)
		if err != nil {
			log.Errorf("Cant send parcel,because %v", err.Error())
			err = m.messageLogic.UpdateStatus(ctx, message.Id, entity.StatusNotSent)
			if err != nil {
				log.Errorf("Cant update message status,because %v", err.Error())
				continue
			}
			continue
		}
		err = m.messageLogic.UpdateStatus(ctx, message.Id, entity.StatusSent)
		if err != nil {
			log.Errorf("Cant update message status,because %v", err.Error())
			continue
		}
	}
}
