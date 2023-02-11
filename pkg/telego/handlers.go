package telego

import (
	"fmt"
	"github.com/SakoDroid/telego/objects"
	objs "github.com/SakoDroid/telego/objects"
)

func (t *telegoClient) registerHandlers() {
	t.startHandler()
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
func (t *telegoClient) messageHandler() chan *objs.Update {
	//Register the channel
	messageChannel, _ := t.bot.AdvancedMode().RegisterChannel("", messageMediaType)
	return *messageChannel
}

func (t *telegoClient) registerChannel() {

}
