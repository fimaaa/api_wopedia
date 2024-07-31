package createeventtestimonyv1

import (
	"backend_base_app/domain/entity"
	"backend_base_app/shared/dbhelpers"
	"context"
)

type apibaseappeventtestimonycreateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &apibaseappeventtestimonycreateInteractor{
		outport: outputPort,
	}
}

func (r *apibaseappeventtestimonycreateInteractor) Execute(ctx context.Context, req entity.EventTestimonyData, testimonyQty int) (*entity.EventTestimonyData, error) {
	err := dbhelpers.WithoutTransaction(ctx, r.outport, func(ctx context.Context) error {
		err := r.outport.CreateEventTestimonyData(ctx, req, testimonyQty)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &req, nil
}
