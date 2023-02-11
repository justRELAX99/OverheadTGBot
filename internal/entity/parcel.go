package entity

type Parcels []Parcel

func (p Parcels) GetMessages() []Message {
	messages := make([]Message, len(p))
	for i, parcel := range p {
		messages[i] = parcel.Message
	}
	return messages
}

func (p Parcels) GetUsers() []User {
	users := make([]User, len(p))
	for i, parcel := range p {
		users[i] = parcel.Sender
	}
	return users
}

type Parcel struct {
	Message Message
	Sender  User
}

type ParcelReceiver interface {
	ReceiverParcels()
}

type MessageSender interface {
	SendMessage(Message)
}
