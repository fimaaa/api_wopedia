package deleteonerelationuserwov1

import (
	"context"
)

type Inport interface {
	Execute(ctx context.Context, id string) (bool, error)
}
