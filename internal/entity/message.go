package entity

import "context"

const (
	StatusReady     = "ready"
	StatusSend      = "sent"
	StatusModerated = "moderated"
	StatusDeleted   = "deleted"
)

type Message struct {
	Text string
	Date int
}

type MessageLogic interface {
	SaveMessagesForModerate(context.Context, []Message) error
	SaveMessageForModerate(context.Context, Message) error
}

type MessageRepository interface {
	SaveMessageForModerate(context.Context, Message) error
}
