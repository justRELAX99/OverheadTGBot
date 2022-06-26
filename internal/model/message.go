package model

const (
	StatusReady     = "ready"
	StatusSend      = "sent"
	StatusModerated = "moderated"
	StatusDeleted   = "deleted"
)

type Message struct {
	Text string
}

type MessageLogic interface {
	SaveMessage([]Message) error
}

type MessageRepository interface {
	SaveMessage([]Message) error
}
