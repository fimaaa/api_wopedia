package createeventguestv1

import (
	"backend_base_app/domain/entity"
	"context"
)

type Inport interface {
	Execute(ctx context.Context, req entity.EventGuestData) (*entity.EventGuestData, error)
}
