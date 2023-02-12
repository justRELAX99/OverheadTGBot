package telego

import (
	"OverheadTGBot/internal/entity"
	"OverheadTGBot/pkg/errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (t telegoClient) HandleCommand(command string) chan entity.Parcel {
	parcelsChannel := make(chan entity.Parcel)
	commandChannel := t.initCommandChannel(command)
	go func() {
		for telegoData := range commandChannel {
			parcelsChannel <- t.convertToParcel(telegoData, command)
		}
	}()
	return parcelsChannel
}

func (t telegoClient) SendToAdminChannel(text string) error {
	return t.SendMessage(t.config.RecipientChatId, text)
}

func (t telegoClient) SendMessage(chatId int64, text string) error {
	msg := tgbotapi.NewMessage(chatId, text)
	_, err := t.bot.Send(msg)
	if err != nil {
		return errors.Wrap(err, "cant send message")
	}
	return nil
}
