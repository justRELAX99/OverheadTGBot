package telego

import (
	"OverheadTGBot/internal/entity"
	"OverheadTGBot/pkg"
	config "OverheadTGBot/pkg/config/entity"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
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

func (t *telegoClient) initCommandChannel(command string) chan tgbotapi.Update {
	commandChannel := make(chan tgbotapi.Update)
	t.commands[command] = commandChannel
	return commandChannel
}

//main channel, who got every message
//can be only one
func (t telegoClient) initUpdatesChannel() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := t.bot.GetUpdatesChan(u)
	if pkg.IsDev() {
		t.bot.Debug = true
	}

	go func() {
		for update := range updates {
			if update.Message == nil { // ignore any non-Message updates
				continue
			}
			if update.Message.IsCommand() {
				t.handleCommand(update)
				continue
			}
			t.messages <- update
		}
	}()

	return
}

func (t telegoClient) handleCommand(update tgbotapi.Update) {
	for k, v := range t.commands {
		if k == update.Message.Command() {
			v <- update
		}
	}
}
