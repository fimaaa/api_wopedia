package deleteoneeventweddingv1

import (
	"context"
)

type Inport interface {
	Execute(ctx context.Context, id string) (bool, error)
}
