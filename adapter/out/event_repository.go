package out

import (
	"context"
	"message-delivery/application/domain"
)

type EventRepository struct{}

func (e EventRepository) Create(ctx context.Context, message *domain.Event) error {
	//TODO implement me
	panic("implement me")
}

func (e EventRepository) GetByID(ctx context.Context, id string) (domain.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (e EventRepository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
