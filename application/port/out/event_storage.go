package out

import (
	"context"
	"message-delivery/application/domain"
)

type EventStorage interface {
	Create(ctx context.Context, message domain.Message) error
	GetByID(ctx context.Context, id string) (domain.Message, error)
	Delete(ctx context.Context, id string) error
}
