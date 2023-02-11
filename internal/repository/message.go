package repository

import (
	"OverheadTGBot/internal/entity"
	"OverheadTGBot/pkg/errors"
	"context"
	"fmt"
)

type messageRepository struct {
	repositoryClient entity.RepositoryClient
}

func NewMessageRepository(repositoryClient entity.RepositoryClient) entity.MessageRepository {
	return &messageRepository{
		repositoryClient: repositoryClient,
	}
}

func (m messageRepository) SaveMessageForModerate(ctx context.Context, message entity.Message) error {
	query := fmt.Sprintf("insert into %s (%s,%s,%s) values(?,?,?)",
		messageTable,
		textColumn,
		dateColumn,
		statusColumn,
	)

	session := m.repositoryClient.GetSession()
	_, err := session.InsertBySql(query, message.Text, message.Date, entity.StatusModerated).ExecContext(ctx)
	if err != nil {
		return errors.Wrap(err, executeError)
	}
	return nil
}
