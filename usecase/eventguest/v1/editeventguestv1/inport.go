package editeventguestv1

import (
	"backend_base_app/domain/entity"
	"context"
)

type Inport interface {
	Execute(ctx context.Context, IDEvent string, req entity.EditEventGuestData) (*entity.EventGuestData, error)
}
