package in

import (
	"context"
	"message-delivery/application/domain"
)

type EventCollector interface {
	Collect(ctx context.Context, message domain.Message) error
}
