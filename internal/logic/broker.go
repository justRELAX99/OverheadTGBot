package logic

import (
	"OverheadTGBot/internal/model"
	"OverheadTGBot/pkg/errors"
	"fmt"
	"log"
	"time"
)

const (
	RedirectTimeout = time.Second * 5
)

type brokerLogic struct {
	parcelRecipient model.ParcelRecipient
	parcelResender  model.ParcelResender
	messageLogic    model.MessageLogic
	userLogic       model.UserLogic
}

func NewBrokerLogic(parcelRecipient model.ParcelRecipient,
	parcelResender model.ParcelResender,
	messageLogic model.MessageLogic,
	userLogic model.UserLogic) model.BrokerLogic {
	return brokerLogic{
		parcelRecipient: parcelRecipient,
		parcelResender:  parcelResender,
		messageLogic:    messageLogic,
		userLogic:       userLogic,
	}
}

func (b brokerLogic) RedirectParcels() {
	for {
		time.Sleep(RedirectTimeout)
		parcels, err := b.getParcels()
		if err != nil {
			log.Println(fmt.Sprintf("Cant redirect parcels,because %v", err.Error()))
		}
		if len(parcels) > 0 {
			err = b.sendParcels(parcels)
			if err != nil {
				log.Println(fmt.Sprintf("Cant redirect parcels,because %v", err.Error()))
			}
		}
	}
}

func (b brokerLogic) getParcels() (parcels []model.Parcel, err error) {
	parcels, err = b.parcelRecipient.GetParcels()
	if err != nil {
		return nil, errors.Wrap(err, "Cant get parcels")
	}
	return parcels, nil
}

func (b brokerLogic) sendParcels(parcels []model.Parcel) (err error) {
	err = b.parcelResender.SendParcels(parcels)
	if err != nil {
		return errors.Wrap(err, "Cant send parcels")
	}
	return nil
}
