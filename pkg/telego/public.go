package telego

import (
	"OverheadTGBot/internal/entity"
	"OverheadTGBot/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (t telegoClient) HandleStart() {
	log := logger.Get()
	startChannel := t.initCommand("start")
	go func() {
		for telegoData := range startChannel {
			msg := tgbotapi.NewMessage(telegoData.Message.Chat.ID, "hi,im telego bot")
			_, err := t.bot.Send(msg)
			if err != nil {
				log.Errorf("cant send message,err = %v", err)
			}
		}
	}()
}

func (t telegoClient) HandleParcels() chan entity.Parcel {
	parcelsChannel := make(chan entity.Parcel)
	go func() {
		for telegoData := range t.messages {

			message := entity.Message{
				Text: telegoData.Message.Text,
				Date: telegoData.Message.Date,
			}

			user := entity.User{
				TelegramId: telegoData.Message.From.ID,
				UserName:   telegoData.Message.From.UserName,
			}
			parcelsChannel <- entity.Parcel{
				Message: message,
				Sender:  user,
			}
		}
	}()
	return parcelsChannel
}

func (t telegoClient) SendMessage(message entity.Message) error {
	return nil
}
