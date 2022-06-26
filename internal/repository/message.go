package repository

import "OverheadTGBot/internal/model"

type messageRepository struct {
	repositoryClient model.RepositoryClient
}

func NewMessageRepository(repositoryClient model.RepositoryClient) model.MessageRepository {
	return &messageRepository{
		repositoryClient: repositoryClient,
	}
}

func (m messageRepository) SaveMessage(messages []model.Message) error {
	//TODO implement me
	panic("implement me")
}
