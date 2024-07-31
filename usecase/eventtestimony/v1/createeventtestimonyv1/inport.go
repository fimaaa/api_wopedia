package createeventtestimonyv1

import (
	"backend_base_app/domain/entity"
	"context"
)

type Inport interface {
	Execute(ctx context.Context, req entity.EventTestimonyData, testimonyQty int) (*entity.EventTestimonyData, error)
}
