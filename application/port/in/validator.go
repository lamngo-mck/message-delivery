package in

import (
	"context"
	"message-delivery/application/domain"
)

type Validator interface {
	Validate(ctx context.Context, message domain.Message) error
}
