package service

import (
	"context"
	"message-delivery/application/domain"
	"message-delivery/application/port/out"
)

type EventCollectorService struct {
	cache      out.Cache
	eventRepo  out.EventRepository
	outboxRepo out.OutboxRepository
	messaging  out.MessageBroker
}

func (e EventCollectorService) Collect(ctx context.Context, event domain.Event) (err error) {
	err = e.persist(ctx, event)
	if err != nil {
		return
	}
	err = e.messaging.Publish(ctx, event)
	return err
}

// TODO: make a transaction for 2 operations in below
func (e EventCollectorService) persist(ctx context.Context, event domain.Event) error {
	err := e.eventRepo.Create(ctx, &event)
	if err != nil {
		return err
	}
	return e.outboxRepo.Create(ctx, &event)
}
