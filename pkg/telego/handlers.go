package telego

import (
	"OverheadTGBot/pkg"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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
