package editeventtestimonyv1

import (
	"backend_base_app/domain/entity"
	"context"
)

type Inport interface {
	Execute(ctx context.Context, req entity.EditEventTestimonyData) (*entity.EventTestimonyData, error)
}
