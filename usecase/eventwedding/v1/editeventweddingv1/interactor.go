package editeventweddingv1

import (
	"backend_base_app/domain/entity"
	"backend_base_app/shared/dbhelpers"
	"context"
)

type apibaseapeventweddingcreateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &apibaseapeventweddingcreateInteractor{
		outport: outputPort,
	}
}

func (r *apibaseapeventweddingcreateInteractor) Execute(ctx context.Context, req entity.EditEventWeddingData) (*entity.EventWeddingData, error) {
	res := &entity.EventWeddingData{}

	err := dbhelpers.WithoutTransaction(ctx, r.outport, func(ctx context.Context) error {
		resUpdate, err := r.outport.UpdateEventWeddingData(ctx, req)
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
