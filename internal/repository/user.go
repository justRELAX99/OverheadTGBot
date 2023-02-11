package repository

import (
	"OverheadTGBot/internal/entity"
	"OverheadTGBot/pkg/errors"
	"context"
	"fmt"
)

type userRepository struct {
	repositoryClient entity.RepositoryClient
}

func NewUserRepository(repositoryClient entity.RepositoryClient) entity.UserRepository {
	return userRepository{
		repositoryClient: repositoryClient,
	}
}

func (u userRepository) SaveUser(ctx context.Context, user entity.User) error {
	query := fmt.Sprintf("insert into %s (%s,%s) values(?,?)",
		userTable,
		telegramIdColumn,
		userNameColumn,
	)

	session := u.repositoryClient.GetSession()
	_, err := session.InsertBySql(query, user.TelegramId, user.UserName).ExecContext(ctx)
	if err != nil {
		return errors.Wrap(err, executeError)
	}
	return nil
}
