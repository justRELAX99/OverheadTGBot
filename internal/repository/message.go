package repository

import "OverheadTGBot/internal/model"

type messageRepository struct {
	repositoryClient model.RepositoryClient
}

func NewMessageRepository(repositoryClietn model.RepositoryClient) model.MessageRepository {
	return &messageRepository{
		repositoryClient: repositoryClietn,
	}
}
