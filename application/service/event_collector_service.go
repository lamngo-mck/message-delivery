package service

import (
	"message-delivery/application/domain"
	"message-delivery/application/port/out"
	"time"
)

type EventCollectorService struct {
	cache out.Cache
}

func (e EventCollectorService) Collect(event domain.Event) error {
	return e.cache.Set("key", event, time.Minute)
}
