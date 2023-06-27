package out

import "context"

type OutboxRepository struct{}

func (OutboxRepository) Create(ctx context.Context, payload *interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (OutboxRepository) GetByID(ctx context.Context, id string) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (OutboxRepository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
