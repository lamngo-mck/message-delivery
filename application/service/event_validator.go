package service

import (
	"context"
	"errors"
	"message-delivery/application/domain"
)

type EventValidator struct{}

func (v EventValidator) Validate(ctx context.Context, message domain.Message) error {
	if len(message.Endpoints) == 0 {
		return errors.New("empty endpoints")
	}
	if message.Schedule != nil && !message.Schedule.IsValid() {
		return errors.New("invalid schedule")
	}
	return nil
}
