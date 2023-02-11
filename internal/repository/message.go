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

func (m messageRepository) SaveMessage(ctx context.Context, message entity.Message) (messageId int, err error) {
	query := fmt.Sprintf("insert into %s (%s,%s,%s) values(?,?,?) returning id;",
		messageTable,
		textColumn,
		dateColumn,
		statusColumn,
	)
	session := m.repositoryClient.GetSession()
	err = session.InsertBySql(query, message.Text, message.Date, message.Status).LoadContext(ctx, &messageId)
	if err != nil {
		return messageId, errors.Wrap(err, executeError)
	}
	return messageId, nil
}

func (m messageRepository) UpdateStatus(ctx context.Context, messageId int, status string) (err error) {
	query := fmt.Sprintf("update %s set %s=? where %s=?",
		messageTable,
		statusColumn,
		messageIdColumn,
	)
	session := m.repositoryClient.GetSession()
	_, err = session.UpdateBySql(query, status, messageId).ExecContext(ctx)
	if err != nil {
		return errors.Wrap(err, executeError)
	}
	return nil
}
