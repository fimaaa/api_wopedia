package editeventtestimonyv1

import (
	"backend_base_app/domain/entity"
	"backend_base_app/shared/dbhelpers"
	"context"
)

type apibaseappediteventtestimonycreateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &apibaseappediteventtestimonycreateInteractor{
		outport: outputPort,
	}
}

func (r *apibaseappediteventtestimonycreateInteractor) Execute(ctx context.Context, req entity.EditEventTestimonyData) (*entity.EventTestimonyData, error) {
	res := &entity.EventTestimonyData{}

	err := dbhelpers.WithoutTransaction(ctx, r.outport, func(ctx context.Context) error {
		resUpdate, err := r.outport.UpdateEventTestimonyData(ctx, req)
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
