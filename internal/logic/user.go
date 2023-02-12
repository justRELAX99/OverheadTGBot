package logic

import (
	"OverheadTGBot/internal/entity"
	"OverheadTGBot/pkg/errors"
	"context"
)

type userLogic struct {
	userRepository entity.UserRepository
}

func NewUserLogic(userRepository entity.UserRepository) entity.UserLogic {
	return userLogic{
		userRepository: userRepository,
	}
}

func (u userLogic) SaveUser(ctx context.Context, user entity.User) error {
	err := u.userRepository.SaveUser(ctx, user)
	if err != nil {
		return errors.Wrap(err, "cant save user")
	}
	return nil
}
