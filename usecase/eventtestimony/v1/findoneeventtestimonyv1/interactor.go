package findoneeventtestimonyv1

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

func (r *apibaseappeventtestimonycreateInteractor) Execute(ctx context.Context, id string) (*entity.EventTestimonyData, error) {
	res := &entity.EventTestimonyData{}

	err := dbhelpers.WithoutTransaction(ctx, r.outport, func(ctx context.Context) error {
		organizerWeddingDataData, err := r.outport.FindOneEventTestimonyDataById(ctx, id)
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
