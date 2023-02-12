package entity

import "context"

type User struct {
	Id         int64
	TelegramId int64
	UserName   string
	Role       Role
}

type UserLogic interface {
	SaveUser(context.Context, User) error
}

type UserRepository interface {
	SaveUser(context.Context, User) error
}
