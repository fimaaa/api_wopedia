package findoneeventguestv1

import (
	"backend_base_app/domain/entity"
	"backend_base_app/shared/dbhelpers"
	"context"
)

type apibaseappeventguestcreateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &apibaseappeventguestcreateInteractor{
		outport: outputPort,
	}
}

func (r *apibaseappeventguestcreateInteractor) Execute(ctx context.Context, id string) (*entity.EventGuestData, error) {
	res := &entity.EventGuestData{}

	err := dbhelpers.WithoutTransaction(ctx, r.outport, func(ctx context.Context) error {
		organizerWeddingDataData, err := r.outport.FindOneEventGuestDataById(ctx, id)
		if err != nil {
			return err
		}

		res = organizerWeddingDataData

		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
