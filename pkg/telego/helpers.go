package telego

import (
	"OverheadTGBot/internal/entity"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

func (t telegoClient) convertToParcel(telegoData tgbotapi.Update, command string) entity.Parcel {
	messageText := strings.Replace(telegoData.Message.Text, fmt.Sprintf("/%s ", command), "", 1)
	message := entity.Message{
		Text: messageText,
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
