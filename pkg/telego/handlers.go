package telego

import (
	"fmt"
	"github.com/SakoDroid/telego/objects"
	"github.com/nikepan/go-datastructures/queue"
	"time"
)

func (t *telegoClient) registerHandlers() {
	t.startHandler()
	t.messageHandler(botTimeout)
}

//Sends the message to the chat that the message has been received from.
//The message will be a reply to the received message.
func (t *telegoClient) startHandler() {
	err := t.bot.AddHandler("/start", func(u *objects.Update) {
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

//Tries to get messages from telegram bot
func (t *telegoClient) messageHandler(timeout time.Duration) {
	//Register the channel
	t.Queue = queue.New(queueSize)

	messageChannel, _ := t.bot.AdvancedMode().RegisterChannel("", messageMediaType)

	go func() {
		for {
			time.Sleep(timeout)
			//Wait for updates
			up := <-*messageChannel
			t.Queue.Put(up)
		}
	}()
}
