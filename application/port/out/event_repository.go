package out

import (
	"context"
	"message-delivery/application/domain"
)

type EventRepository interface {
	Create(ctx context.Context, message *domain.Event) error
	GetByID(ctx context.Context, id string) (domain.Event, error)
	Delete(ctx context.Context, id string) error
}
