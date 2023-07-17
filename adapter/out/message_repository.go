package out

import (
	"context"
	"database/sql"
	"message-delivery/application/domain"
)

type MessageRepository struct {
	*sql.DB
}

func (e MessageRepository) Create(ctx context.Context, message *domain.Message) error {
	//TODO implement me
	panic("implement me")
}

func (e MessageRepository) GetByID(ctx context.Context, id string) (domain.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (e MessageRepository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
