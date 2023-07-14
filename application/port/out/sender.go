package out

import "context"

type Sender interface {
	Send(ctx context.Context, content string, endpoint []string) error
}
