package model

type Parcel struct {
	Message Message
	Sender  User
}

type ParcelRecipient interface {
	GetParcels() ([]Parcel, error)
}

type ParcelResender interface {
	SendParcels([]Parcel) error
}
