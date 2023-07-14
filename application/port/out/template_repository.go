package out

import (
	"context"
)

type TemplateStorage interface {
	GetByID(ctx context.Context, id string) (string, error)
}
