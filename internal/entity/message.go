package entity

import "context"

const (
	StatusReady     = "ready"
	StatusSent      = "sent"
	StatusNotSent   = "not_sent"
	StatusModerated = "moderated"
	StatusDeleted   = "deleted"
)

type Message struct {
	Id     int
	Text   string
	Date   int
	Status string
}

type MessageLogic interface {
	SaveMessage(context.Context, Message) (int, error)
	UpdateStatus(context.Context, int, string) error
}

type MessageRepository interface {
	SaveMessage(context.Context, Message) (int, error)
	UpdateStatus(context.Context, int, string) error
}
