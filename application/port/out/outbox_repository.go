package out

import "context"

type OutboxRepository interface {
	Create(ctx context.Context, payload interface{}) error
	GetByID(ctx context.Context, id string) (interface{}, error)
	Delete(ctx context.Context, id string) error
}
