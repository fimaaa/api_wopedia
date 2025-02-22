package createeventweddingv1

import (
	"backend_base_app/domain/entity"
	"context"
)

type Inport interface {
	Execute(ctx context.Context, req entity.EventWeddingData) (*entity.EventWeddingData, error)
}
