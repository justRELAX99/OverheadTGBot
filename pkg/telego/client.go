package telego

import (
	"OverheadTGBot/internal/entity"
	config "OverheadTGBot/pkg/config/entity"
	"OverheadTGBot/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"

	"log"
)

//Chat types
const (
	privateChatType    = "private"
	groupChatType      = "group"
	superGroupChatType = "supergroup"
	channelChatType    = "channel"
	allChatType        = "all"
)

const (
	messageMediaType = "message"

	//number of messages we are trying to get from the queue
	countMessage = 100
	//time after which we try to receive messages from the telegram bot
	botTimeout = time.Second * 5
	//queue size
	queueSize    = 1000
	queueTimeout = time.Second * 5
)

type telegoClient struct {
	config config.TelegramBotConfig
	bot    *tgbotapi.BotAPI

	commands map[string]chan tgbotapi.Update
	messages chan tgbotapi.Update
}

func NewTelegoClient(config config.TelegramBotConfig) entity.TelegramClient {
	client := telegoClient{
		config:   config,
		commands: make(map[string]chan tgbotapi.Update),
		messages: make(chan tgbotapi.Update),
	}
	client.initClient()
	client.HandleStart()
	client.initUpdatesChannel()
	return client
}

func (t *telegoClient) initClient() {
	bot, err := tgbotapi.NewBotAPI(t.config.HttpToken)
	if err != nil {
		log.Fatal(err)
	}
	t.bot = bot
	return
}

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

func (t *telegoClient) initCommand(command string) chan tgbotapi.Update {
	commandChannel := make(chan tgbotapi.Update)
	t.commands[command] = commandChannel
	return commandChannel
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
