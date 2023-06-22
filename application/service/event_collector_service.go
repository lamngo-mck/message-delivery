package service

import "message-delivery/application/domain"

type EventCollectorService struct {
}

func (e EventCollectorService) Collect(event domain.Event) error {
	return nil
}
