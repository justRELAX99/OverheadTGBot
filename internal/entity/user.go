package entity

import "context"

type User struct {
	Id         int64
	TelegramId int64
	UserName   string
}

type UserLogic interface {
	SaveUsers(context.Context, []User) error
	SaveUser(context.Context, User) error
}

type UserRepository interface {
	SaveUser(context.Context, User) error
}
