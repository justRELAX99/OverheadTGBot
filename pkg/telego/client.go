package telego

import (
	"OverheadTGBot/internal/entity"
	config "OverheadTGBot/pkg/config/entity"
	"OverheadTGBot/pkg/errors"
	"github.com/SakoDroid/telego"
	configTelego "github.com/SakoDroid/telego/configs"
	"time"

	"github.com/nikepan/go-datastructures/queue"
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
	config      config.TelegramBotConfig
	bot         *telego.Bot
	downTimeout int
	Queue       *queue.Queue
}

func NewTelegoClient(config config.TelegramBotConfig) entity.TelegramClient {
	client := initClient(config)
	client.registerHandlers()
	return client
}

func initClient(config config.TelegramBotConfig) telegoClient {
	//Bot configs
	cf := configTelego.BotConfigs{
		BotAPI:         configTelego.DefaultBotAPI,
		APIKey:         config.HttpToken,
		UpdateConfigs:  configTelego.DefaultUpdateConfigs(),
		Webhook:        false,
		LogFileAddress: configTelego.DefaultLogFile,
	}

	bot, err := telego.NewBot(&cf)
	if err != nil {
		log.Fatal(err)
	}
	//Start the bot.
	err = bot.Run()
	if err != nil {
		log.Fatal(err)
	}
	client := telegoClient{
		config: config,
		bot:    bot,
	}
	return client
}

func (t telegoClient) HandleParcels() chan entity.Parcel {
	parcelsChannel := make(chan entity.Parcel)
	telegoMessageChannel := t.messageHandler()
	go func() {
		for telegoData := range telegoMessageChannel {

			message := entity.Message{
				Text: telegoData.Message.Text,
				Date: telegoData.Message.Date,
			}

			user := entity.User{
				TelegramId: telegoData.Message.From.Id,
				UserName:   telegoData.Message.From.Username,
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
	_, err := t.bot.SendMessageUN(t.config.RecipientChatId, message.Text+" resend", "", 0, false, false)
	if err != nil {
		return errors.Wrap(err, "cant send message")
	}
	return nil
}
