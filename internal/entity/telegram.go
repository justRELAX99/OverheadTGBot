package entity

type TelegramClient interface {
	HandleParcels() chan Parcel
	SendParcels(Parcels) error
}
