package telego

import (
	"OverheadTGBot/internal/model"
	config "OverheadTGBot/pkg/config/model"
	"OverheadTGBot/pkg/errors"
	"github.com/SakoDroid/telego"
	configTelego "github.com/SakoDroid/telego/configs"
	objs "github.com/SakoDroid/telego/objects"
	"time"

	"github.com/nikepan/go-datastructures/queue"
	"log"
	"strings"
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

func NewTelegoClient(config config.TelegramBotConfig) telegoClient {
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

func (t telegoClient) GetParcels() (parcels []model.Parcel, err error) {
	var datas []interface{}
	datas, err = t.Queue.Poll(countMessage, queueTimeout)
	if err != nil {
		if strings.Contains(err.Error(), "queue: poll timed out") {
			return nil, nil
		}
		return nil, errors.Wrap(err, "Cant get messages from pool")
	}

	for _, data := range datas {
		if telegoData, ok := data.(*objs.Update); ok {
			message := model.Message{
				Text: telegoData.Message.Text,
			}
			user := model.User{
				UserName: telegoData.Message.From.Username,
			}
			parcels = append(parcels, model.Parcel{
				Message: message,
				Sender:  user,
			})
		}
	}
	return parcels, nil
}

func (t telegoClient) SendParcels([]model.Parcel) error {
	return nil
}
