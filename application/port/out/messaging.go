package out

import "context"

type MessageBroker interface {
	Publish(ctx context.Context, event interface{}) error
	Process(ctx context.Context) error
}
