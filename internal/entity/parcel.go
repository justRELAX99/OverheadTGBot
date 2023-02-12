package entity

type MessageSender interface {
	SendMessage(Message)
}
