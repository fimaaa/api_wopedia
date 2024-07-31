package editeventweddingv1

import (
	"backend_base_app/domain/entity"
	"context"
)

type Inport interface {
	Execute(ctx context.Context, req entity.EditEventWeddingData) (*entity.EventWeddingData, error)
}
