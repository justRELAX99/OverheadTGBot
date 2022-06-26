package logic

import "OverheadTGBot/internal/model"

type userLogic struct {
	userRepository model.UserRepository
}

func NewUserLogic(userRepository model.UserRepository) model.UserLogic {
	return userLogic{
		userRepository: userRepository,
	}
}
