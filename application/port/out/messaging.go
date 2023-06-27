package out

import "context"

type MessageBroker interface {
	Publish(ctx context.Context, event interface{}) error
	Consume(ctx context.Context) error
}
