package repository

import "OverheadTGBot/internal/model"

type userRepository struct {
	repositoryClient model.RepositoryClient
}

func NewUserRepository(repositoryClient model.RepositoryClient) model.UserRepository {
	return userRepository{
		repositoryClient: repositoryClient,
	}
}
