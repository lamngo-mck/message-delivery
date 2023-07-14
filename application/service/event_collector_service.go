package service

import (
	"context"
	"message-delivery/application/domain"
	"message-delivery/application/port/out"
	"time"
)

type EventCollectorService struct {
	messageRepo  out.MessageRepository
	eventStorage out.EventStorage
	messaging    out.MessageBroker
}

func NewEventCollectorService() EventCollectorService {
	return EventCollectorService{}
}

func (e EventCollectorService) Collect(ctx context.Context, message domain.Message) (err error) {
	err = e.messageRepo.Create(ctx, &message)
	if err != nil {
		return
	}
	if message.IsBatching {
		err = e.eventStorage.Create(ctx, message)
		return
	}
	event := domain.Event{
		ID:        message.ID,
		Timestamp: time.Now().UTC().Unix(),
	}
	err = e.messaging.Publish(ctx, event)
	return err
}
