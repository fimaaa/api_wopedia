package editeventguestv1

import (
	"backend_base_app/domain/entity"
	"backend_base_app/shared/dbhelpers"
	"context"
)

type apibaseapeventguestcreateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &apibaseapeventguestcreateInteractor{
		outport: outputPort,
	}
}

func (r *apibaseapeventguestcreateInteractor) Execute(ctx context.Context, IDEvent string, req entity.EditEventGuestData) (*entity.EventGuestData, error) {
	res := &entity.EventGuestData{}

	err := dbhelpers.WithoutTransaction(ctx, r.outport, func(ctx context.Context) error {
		resUpdate, err := r.outport.UpdateEventGuestData(ctx, IDEvent, req)
		if err != nil {
			return err
		}
		res = resUpdate

		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
