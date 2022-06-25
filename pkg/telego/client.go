package telego

import (
	"OverheadTGBot/internal/model"
	config "OverheadTGBot/pkg/config/model"
	"fmt"
	"github.com/SakoDroid/telego"
	configTelego "github.com/SakoDroid/telego/configs"
	"github.com/SakoDroid/telego/objects"
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
)

type telegoClient struct {
	config config.TelegramBotConfig
	bot    *telego.Bot
}

func NewClient(config config.TelegramBotConfig) model.TelegramBot {
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
	telegoClient := telegoClient{
		config: config,
		bot:    bot,
	}
	telegoClient.startHandler()
	return telegoClient
}

func (t telegoClient) startHandler() {
	err := t.bot.AddHandler("/start", func(u *objects.Update) {
		//Sends the message to the chat that the message has been received from.
		//The message will be a reply to the received message.
		_, err := t.bot.SendMessage(
			u.Message.Chat.Id, "hi i'm a telegram bot!",
			"",
			u.Message.MessageId,
			false,
			false)
		if err != nil {
			fmt.Println(err)
		}

	}, privateChatType, groupChatType)
	if err != nil {
		return
	}
}

func (t telegoClient) RegisterMessageHandler() {
	//Register the channel
	messageChannel, _ := t.bot.AdvancedMode().RegisterChannel("", messageMediaType)

	for {
		//Wait for updates
		up := <-*messageChannel

		//Print the text
		fmt.Println(up.Message.Text)
	}
}
