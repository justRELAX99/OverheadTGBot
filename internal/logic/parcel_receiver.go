package logic

import (
	"OverheadTGBot/internal/entity"
	"OverheadTGBot/pkg/logger"
	"context"
)

type parcelReceiver struct {
	parcelChannel chan entity.Parcel
	messageLogic  entity.MessageLogic
	userLogic     entity.UserLogic
}

func NewParcelReceiver(
	telegramClient entity.TelegramClient,
	messageLogic entity.MessageLogic,
	userLogic entity.UserLogic,
) entity.ParcelReceiver {
	return parcelReceiver{
		parcelChannel: telegramClient.HandleParcels(),
		messageLogic:  messageLogic,
		userLogic:     userLogic,
	}
}

func (p parcelReceiver) ReceiverParcels() {
	go p.receiveParcels()
}

func (p parcelReceiver) receiveParcels() {
	ctx := context.Background()
	log := logger.Get()

	for parcel := range p.parcelChannel {
		err := p.messageLogic.SaveMessageForModerate(ctx, parcel.Message)
		if err != nil {
			log.Errorf("Cant save messages,because %v", err.Error())
		}

		err = p.userLogic.SaveUser(ctx, parcel.Sender)
		if err != nil {
			log.Errorf("Cant save users,because %v", err.Error())
		}
	}
}
