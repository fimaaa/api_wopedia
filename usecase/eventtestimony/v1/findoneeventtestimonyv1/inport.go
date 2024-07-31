package findoneeventtestimonyv1

import (
	"backend_base_app/domain/entity"
	"context"
)

type Inport interface {
	Execute(ctx context.Context, id string) (*entity.EventTestimonyData, error)
}
