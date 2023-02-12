package telego

import (
	"OverheadTGBot/internal/entity"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (t telegoClient) convertToParcel(telegoData tgbotapi.Update) entity.Parcel {
	message := entity.Message{
		Text: telegoData.Message.Text,
		Date: telegoData.Message.Date,
	}

	user := entity.User{
		TelegramId: telegoData.Message.From.ID,
		UserName:   telegoData.Message.From.UserName,
	}

	return entity.Parcel{
		Message: message,
		Sender:  user,
	}
}
