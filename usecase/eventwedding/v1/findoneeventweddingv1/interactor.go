package findoneeventweddingv1

import (
	"backend_base_app/domain/entity"
	"backend_base_app/shared/dbhelpers"
	"context"
)

type apibaseappeventweddingcreateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &apibaseappeventweddingcreateInteractor{
		outport: outputPort,
	}
}

func (r *apibaseappeventweddingcreateInteractor) Execute(ctx context.Context, id string) (*entity.EventWeddingData, error) {
	res := &entity.EventWeddingData{}

	err := dbhelpers.WithoutTransaction(ctx, r.outport, func(ctx context.Context) error {
		organizerWeddingDataData, err := r.outport.FindOneEventWeddingDataById(ctx, id)
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
