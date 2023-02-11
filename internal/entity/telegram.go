package entity

type TelegramClient interface {
	HandleParcels() chan Parcel
	SendMessage(Message) error
}
